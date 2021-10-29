package main

// sample is a timeseries sample with a timestamp and value
type sample struct {
	ts  int64
	val float64
}

// Series holds samples for a timeseries.
// You may add samples in order or out of order
type Series interface {
	Add(s sample)
	Iterator() Iterator
}

// Iterator returns samples in timestamp ascending order
type Iterator interface {
	Next() bool
	At() sample
}
