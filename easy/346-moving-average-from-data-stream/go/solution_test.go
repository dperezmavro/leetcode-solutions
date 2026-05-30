package moving_average_from_data_stream

import (
	"fmt"
	"math"
	"testing"
)

func TestMovingAverage(t *testing.T) {
	tests := []struct {
		input   []int
		k       int
		outputs []float64
	}{
		{
			k:       3,
			input:   []int{1, 10, 3, 5},
			outputs: []float64{1.0, 5.5, 4.66667, 6.0},
		},

		{
			k: 5,
			input: []int{
				12009,
				1965,
				-940,
				-8516,
				-16446,
				7870,
				25545,
				-21028,
				18430,
				-23464,
			},
			outputs: []float64{
				12009.00000,
				6987.00000,
				4344.66667,
				1129.50000,
				-2385.60000,
				-3213.40000,
				1502.60000,
				-2515.00000,
				2874.20000,
				1470.60000,
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.input), func(t *testing.T) {

			mv := Constructor(tt.k)
			for i, v := range tt.input {
				out := mv.Next(v)
				if math.Abs(tt.outputs[i]-out) > 0.5 {
					t.Errorf("error in output %d: wanted %f, got %f", i, tt.outputs[i], out)
				}
			}
		})
	}
}
