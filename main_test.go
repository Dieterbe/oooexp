package main

import (
	"math/rand"
	"reflect"
	"testing"
)

// shuffle shuffles factor-sized subgroups within s
func shuffle(s []sample, factor int) {

	if len(s)%factor != 0 {
		panic("shuffling method isn't smart enough to deal with this. don't do this please")
	}

	// switch the first and last and mix up the points between them if there's 2 or more.
	for i, j := 0, factor; j <= len(s); i, j = j, j+factor {
		rand.Shuffle(factor, func(x, y int) {
			s[x+i], s[y+i] = s[y+i], s[x+i]
		})
	}
}

func TestIngestOOO(t *testing.T) {

	// here we have 20 items, out of ordere in a factor of 5
	// in reality we may have hundreds of items, out of order in a factor from 20 to 200
	var input []sample
	for i := 0; i < 20; i++ {
		input = append(input, sample{ts: int64(i), val: float64(i)})
	}
	sortedInput := make([]sample, len(input))
	copy(sortedInput, input)
	shuffle(input, 5)

	series := &rawSeriesSortOnQuery{}
	for _, sample := range input {
		series.Add(sample)
	}
	it := series.Iterator()

	var got []sample
	for it.Next() {
		got = append(got, it.At())
	}

	if !reflect.DeepEqual(sortedInput, got) {
		t.Fatalf("not equal\nexp:%v\ngot:%v", sortedInput, got)
	}

}
