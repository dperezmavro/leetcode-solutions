package main

import (
	"bytes"
	"sync"
	"testing"
)

func TestMain(t *testing.T) {

	tests := []struct {
		n    int
		want string
	}{
		{
			n:    1,
			want: "foobar",
		},
		{
			n:    2,
			want: "foobarfoobar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {

			var buf bytes.Buffer

			fb := NewFooBar(tt.n)

			wg := sync.WaitGroup{}
			wg.Add(2)

			wg.Go(func() {
				fb.Foo(func() {
					buf.Write([]byte("foo"))
				})
				wg.Done()
			})
			wg.Go(func() {
				fb.Bar(func() {
					buf.Write([]byte("bar"))
				})
				wg.Done()
			})

			wg.Wait()

			if buf.String() != tt.want {
				t.Errorf("wrong output: wanted %s, got %s", tt.want, buf.String())
			}
		})
	}
}
