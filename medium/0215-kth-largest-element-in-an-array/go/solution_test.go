package kleia

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		input  []int
		result int
		k      int
	}{
		{
			input:  []int{3, 2, 1, 5, 6, 4},
			k:      2,
			result: 5,
		},
		{
			input:  []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:      4,
			result: 4,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v_%d_%d", tt.input, tt.k, tt.result), func(t *testing.T) {
			result := findKthLargest(tt.input, tt.k)
			if result != tt.result {
				t.Errorf("wanted %d, got %d", tt.result, result)
			}
		})
	}
}
