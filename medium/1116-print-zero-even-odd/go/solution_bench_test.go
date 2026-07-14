package print_zero_even_odd

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// package-level sink prevents the compiler from eliding printNumber calls
var sink int

// safe without atomics: printNumber is always invoked while holding z.cond.L
func runOnce(n int) {
	zeo := NewZeroEvenOdd(n)
	record := func(i int) { sink += i }

	wg := sync.WaitGroup{}
	wg.Go(func() { zeo.Zero(record) })
	wg.Go(func() { zeo.Even(record) })
	wg.Go(func() { zeo.Odd(record) })
	wg.Wait()
}

func BenchmarkZeroEvenOdd(b *testing.B) {
	for _, n := range []int{10, 100, 1000} {
		for _, procs := range []int{1, runtime.NumCPU()} {
			b.Run(fmt.Sprintf("n=%d/procs=%d", n, procs), func(b *testing.B) {
				prev := runtime.GOMAXPROCS(procs)
				defer runtime.GOMAXPROCS(prev)

				b.ReportAllocs()
				for b.Loop() {
					runOnce(n)
				}
			})
		}
	}
}
