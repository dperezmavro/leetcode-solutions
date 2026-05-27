package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwitter(t *testing.T) {
	tests := []struct {
		input []int
		k     int
		out   []int
	}{
		{
			input: []int{1, 1, 1, 2, 2, 3},
			k:     2,
			out:   []int{1, 2},
		},
		{
			input: []int{1},
			k:     1,
			out:   []int{1},
		},
		{
			input: []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2},
			k:     2,
			out:   []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v - %d", tt.input, tt.k), func(t *testing.T) {
			res := topKFrequent(tt.input, tt.k)
			assert.Equal(t, tt.out, res)
		})
	}
}
