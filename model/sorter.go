package model

import (
	"sort"
)

type pathMapSorter struct {
	paths []PathMap
	by    sortBy
}

type sortBy func(p1, p2 *PathMap) bool

func (r *Redirect) sortPathMap() {
	var fr sortBy
	if r.PathMaps == nil {
		return
	}
	fr = func(p1, p2 *PathMap) bool {
		return len(p1.From) > len(p2.From)
	}
	fr.Sort(r.PathMaps)
}

func (by sortBy) Sort(paths []PathMap) {
	ps := &pathMapSorter{
		paths: paths,
		by:    by,
	}
	sort.Sort(ps)
}

func (s *pathMapSorter) Len() int {
	return len(s.paths)
}

func (s *pathMapSorter) Swap(i, j int) {
	s.paths[i], s.paths[j] = s.paths[j], s.paths[i]
}

func (s *pathMapSorter) Less(i, j int) bool {
	return s.by(&s.paths[i], &s.paths[j])
}
