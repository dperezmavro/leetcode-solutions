package simplifypath

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 1},
			output: 1,
		},
		{
			input:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			output: 49,
		},
		{
			input:  []int{1, 2},
			output: 1,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.input), func(t *testing.T) {
			result := maxArea(tt.input)
			if tt.output != result {
				t.Errorf("error: wanted %d, got %d", tt.output, result)
			}
		})
	}

}
