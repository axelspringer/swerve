package model

import (
	"net/url"
	"path"
	"strings"
)

// GetRedirect returns the calculated route
func (r *Redirect) GetRedirect(reqURL *url.URL) (string, int) {
	code := r.Code
	reURL := r.RedirectTo
	rePath := ""
	reQuery := ""
	reqPath := reqURL.EscapedPath()
	reqQuery := reqURL.RawQuery

	if r.Promotable == true {
		rePath = reqURL.Path

		if len(reqURL.RawQuery) > 0 {
			reQuery = "?" + reqURL.RawQuery
		}
	}

	if r.PathMaps != nil && len(r.PathMaps) > 0 {
		for _, p := range r.PathMaps {
			if p.To == "" {
				continue
			}
			if strings.HasPrefix(reqPath+"?"+reqQuery, p.From) {
				if strings.HasPrefix(reqPath, p.From) {
					rePath = reqPath[len(p.From):]
				} else {
					rePath = p.From
				}
				if strings.HasPrefix(p.To, "http://") || strings.HasPrefix(p.To, "https://") {
					reURL = p.To
				} else {
					if r.Promotable {
						rePath = path.Join(p.To, rePath)
					} else {
						rePath = p.To
					}
				}
				break
			}
		}
	}

	if strings.HasSuffix(reURL, "/") && strings.HasPrefix(rePath, "/") {
		rePath = strings.TrimLeft(rePath, "/")
	}

	return reURL + rePath + reQuery, code
}
