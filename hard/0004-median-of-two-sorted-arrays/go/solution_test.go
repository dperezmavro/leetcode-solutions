package solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		result float64
		input1 []int
		input2 []int
	}{
		{
			input1: []int{1, 3},
			input2: []int{2},
			result: 2.0,
		},
		{
			input1: []int{1, 2},
			input2: []int{3, 4},
			result: 2.5,
		},
		{
			input1: []int{-10, -9, -8},
			input2: []int{1, 2},
			result: -8.0,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%f %+v %+v", tt.result, tt.input1, tt.input2), func(t *testing.T) {
			res := findMedianSortedArrays(tt.input1, tt.input2)
			if res != tt.result {
				t.Errorf("mismatch: wanted %f, got %f", tt.result, res)
			}
		})
	}
}
