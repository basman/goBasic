package main

import (
	"fmt"
	"testing"
)

func TestComputeDigitSum(t *testing.T) {
	// arrange
	tbl := []struct {
		input, outcome int
	}{
		{7, 7},
		{21, 3},
		{101, 2},
		{0101, 2},
		{999, 27},
	}

	for _, tt := range tbl {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			// act
			s := computeDigitSum(tt.input)

			// assert
			if s != tt.outcome {
				t.Errorf("invalid result %v, expected %v", s, tt.outcome)
			}
		})
	}
}
