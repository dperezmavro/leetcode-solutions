package kth_largest_element

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthLargest(t *testing.T) {
	tests := []struct {
		k        int
		starting []int
		ops      []int
		results  []int
	}{
		{
			k:        3,
			starting: []int{4, 5, 8, 2},
			ops:      []int{3, 5, 10, 9, 4},
			results:  []int{4, 5, 5, 8, 8},
		},
		{
			k:        4,
			starting: []int{7, 7, 7, 7, 8, 3},
			ops:      []int{2, 10, 9, 9},
			results:  []int{7, 7, 7, 8},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %+v", tt.k, tt.starting), func(t *testing.T) {
			kthLargest := Constructor(tt.k, tt.starting)
			results := []int{}
			for _, v := range tt.ops {
				results = append(results, kthLargest.Add(v))
			}

			if !assert.Equal(
				t,
				tt.results,
				results,
			) {
				t.Errorf("wanted %+v, got %+v", tt.results, results)
			}
		})
	}
}
