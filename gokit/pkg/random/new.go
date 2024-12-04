package random

import (
	"fmt"
)

type Generator struct {
	seeds []any
	nSeed int
}

func NewGenerator(values []any, rates []uint64) (*Generator, error) {
	lenRate := len(rates) //nolint: ifshort

	if lenValue := len(values); lenValue != lenRate {
		return nil, fmt.Errorf("length of values (%d) and rates (%d) are different", lenValue, lenRate)
	}
	var seeds []any
	sum := uint64(0)

	for idx, rate := range rates {
		for i := uint64(0); i < rate; i++ {
			seeds = append(seeds, values[idx])
		}
		sum += rate
	}

	return &Generator{
		seeds: seeds,
		nSeed: len(seeds),
	}, nil
}

func (r *Generator) Next() any {
	return r.seeds[Int64(0, int64(r.nSeed))]
}
