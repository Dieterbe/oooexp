package main

import "sort"

type rawSeriesSortOnQuery struct {
	recv []sample
}

func (r *rawSeriesSortOnQuery) Add(s sample) {
	r.recv = append(r.recv, s)
}

func (r *rawSeriesSortOnQuery) Iterator() Iterator {
	return newSortedIterator(r.recv)
}

type sortedIterator struct {
	pos  int
	data []sample
}

type byts []sample

func (a byts) Len() int           { return len(a) }
func (a byts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byts) Less(i, j int) bool { return a[i].ts < a[j].ts }

func newSortedIterator(in []sample) *sortedIterator {
	data := make([]sample, len(in))
	copy(data, in)
	sort.Sort(byts(data))
	return &sortedIterator{
		pos:  -1,
		data: data,
	}
}

func (s *sortedIterator) Next() bool {
	if s.pos == len(s.data)-1 {
		return false
	}
	s.pos++
	return true
}

func (s *sortedIterator) At() sample {
	return s.data[s.pos]
}
