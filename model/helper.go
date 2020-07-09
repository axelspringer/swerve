package model

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"

	"github.com/axelspringer/swerve/database"
)

func compress(r Redirect) (database.Redirect, error) {
	var buf bytes.Buffer
	r.sortPathMap()
	compressed := database.Redirect{
		RedirectFrom: r.RedirectFrom,
		Description:  r.Description,
		RedirectTo:   r.RedirectTo,
		Promotable:   r.Promotable,
		Code:         r.Code,
		Created:      r.Created,
		Modified:     r.Modified,
	}

	if r.PathMaps == nil {
		return compressed, nil
	}
	pm, err := json.Marshal(r.PathMaps)
	if err != nil {
		return compressed, errors.WithMessage(err, ErrPathMapsMarshal)
	}
	writer := gzip.NewWriter(&buf)
	if _, err := writer.Write(pm); err != nil {
		return compressed, errors.WithMessage(err, ErrWriterCreate)
	}
	if err := writer.Flush(); err != nil {
		return compressed, errors.WithMessage(err, ErrWriterFlush)
	}

	if err := writer.Close(); err != nil {
		return compressed, errors.WithMessage(err, ErrWriterClose)
	}

	bytes := buf.Bytes()
	compressed.CPathMaps = &bytes
	return compressed, nil
}

func multiCompress(redirects []Redirect) ([]database.Redirect, error) {
	compresseds := []database.Redirect{}
	for _, redirect := range redirects {
		compressed, err := compress(redirect)
		if err != nil {
			return nil, errors.WithMessage(err, ErrRedirectCompress)
		}
		compresseds = append(compresseds, compressed)
	}
	return compresseds, nil
}

func decompress(compressed database.Redirect) (Redirect, error) {
	var pm []PathMap
	decompressed := Redirect{
		RedirectFrom: compressed.RedirectFrom,
		Description:  compressed.Description,
		RedirectTo:   compressed.RedirectTo,
		Promotable:   compressed.Promotable,
		Code:         compressed.Code,
		Created:      compressed.Created,
		Modified:     compressed.Modified,
	}

	if compressed.CPathMaps == nil {
		return decompressed, nil
	}
	buf := bytes.NewBuffer(*compressed.CPathMaps)
	reader, err := gzip.NewReader(buf)
	if err != nil {
		return decompressed, errors.WithMessage(err, ErrReaderCreate)
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return decompressed, errors.WithMessage(err, ErrReaderRead)
	}

	if err := reader.Close(); err != nil {
		return decompressed, errors.WithMessage(err, ErrReaderClose)
	}

	err = json.Unmarshal(data, &pm)
	if err != nil {
		return decompressed, errors.WithMessage(err, ErrPathMapsUnmarshal)
	}

	decompressed.PathMaps = pm
	return decompressed, nil
}

func multiDecompress(compresseds []database.Redirect) ([]Redirect, error) {
	redirects := []Redirect{}
	for _, compressed := range compresseds {
		redirect, err := decompress(compressed)
		if err != nil {
			return nil, errors.WithMessage(err, ErrRedirectsDecompress)
		}
		redirects = append(redirects, redirect)
	}
	return redirects, nil
}
