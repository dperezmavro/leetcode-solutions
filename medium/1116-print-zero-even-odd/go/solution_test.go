package print_zero_even_odd

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestCode(t *testing.T) {

	tests := map[string]struct {
		n    int
		want string
	}{
		"2": {
			n:    2,
			want: "0102",
		},
		"3": {
			n:    3,
			want: "010203",
		},
	}
	for k, tt := range tests {
		t.Run(k, func(t *testing.T) {

			zeo := NewZeroEvenOdd(tt.n)

			wg := sync.WaitGroup{}
			b := bytes.NewBuffer([]byte{})

			wg.Go(func() {
				t.Helper()
				zeo.Zero(func(i int) {
					fmt.Printf("%d", i)
					io.WriteString(b, fmt.Sprintf("%d", i))
				})
			})

			wg.Go(func() {
				t.Helper()
				zeo.Even(func(i int) {
					fmt.Printf("%d", i)
					io.WriteString(b, fmt.Sprintf("%d", i))
				})
			})

			wg.Go(func() {
				t.Helper()
				zeo.Odd(func(i int) {
					fmt.Printf("%d", i)
					io.WriteString(b, fmt.Sprintf("%d", i))
				})
			})

			wg.Wait()

			assert.Equal(t, b.String(), tt.want)
		})
	}
}
