package sliding_window_maximum

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		nums   []int
		output []int
		k      int
	}{
		{
			k:      1,
			nums:   []int{1},
			output: []int{1},
		},
		{
			k:      1,
			nums:   []int{1, -1},
			output: []int{1, -1},
		},
		{
			k:      2,
			nums:   []int{7, 2, 4},
			output: []int{7, 4},
		},
		{
			k:      3,
			nums:   []int{1, 3, -1, -3, 5, 3, 6, 7},
			output: []int{3, 3, 5, 5, 6, 7},
		},
		{
			k:      5,
			nums:   []int{9, 10, 9, -7, -4, -8, 2, -6},
			output: []int{10, 10, 9, 2},
		},
		{
			k:      3,
			nums:   []int{9, 8, 9, 8},
			output: []int{9, 9},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.nums), func(t *testing.T) {
			output := maxSlidingWindow(tt.nums, tt.k)
			assert.Equal(t, output, tt.output)
		})
	}
}
