package asteroidcolissions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsteroidCollisions(t *testing.T) {

	tests := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{5, 10, -5},
			want:  []int{5, 10},
		},
		{
			input: []int{8, -8},
			want:  []int{},
		},
		{
			input: []int{10, 2, -5},
			want:  []int{10},
		},
		{
			input: []int{3, 5, -6, 2, -1, 4},
			want:  []int{-6, 2, 4},
		},
		{
			input: []int{-2, -2, 1, -2},
			want:  []int{-2, -2, -2},
		},
		{
			input: []int{-2, -2, 1, -1},
			want:  []int{-2, -2},
		},
		{
			input: []int{1, -1, -2, -2},
			want:  []int{-2, -2},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.input), func(t *testing.T) {
			res := asteroidCollision(tt.input)
			if !assert.Equal(t, tt.want, res) {
				t.Errorf("wanted %+v, got %+v", tt.want, res)
			}
		})
	}

}
