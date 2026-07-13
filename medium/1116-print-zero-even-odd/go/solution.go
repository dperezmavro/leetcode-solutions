package print_zero_even_odd

import (
	"sync"
)

type ZeroEvenOdd struct {
	n       int
	counter int
	zero    bool
	cond    *sync.Cond
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	zeo := &ZeroEvenOdd{
		n:       n,
		counter: 1,
		zero:    true,
		cond:    sync.NewCond(&sync.Mutex{}),
	}

	return zeo
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for {
		z.cond.L.Lock()

		for !z.zero {
			z.cond.Wait()
		}

		// added this to avoid having to calculate
		// bounds ranges in outter for-loop
		if z.counter > z.n {
			z.cond.L.Unlock()
			return
		}
		printNumber(0)
		z.zero = false

		z.cond.Broadcast()
		z.cond.L.Unlock()
	}

}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for {
		z.cond.L.Lock()

		for z.counter <= z.n && (z.counter%2 == 0 || z.zero) {
			z.cond.Wait()
		}

		// added this to avoid having to calculate
		// bounds ranges in outter for-loop
		if z.counter > z.n {
			z.cond.L.Unlock()
			return
		}

		printNumber(z.counter)
		z.counter += 1
		z.zero = true

		z.cond.Broadcast()
		z.cond.L.Unlock()
	}
}
func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	for {
		z.cond.L.Lock()

		for z.counter <= z.n && (z.counter%2 != 0 || z.zero) {
			z.cond.Wait()
		}

		// added this to avoid having to calculate
		// bounds ranges in outter for-loop
		if z.counter > z.n {
			z.cond.L.Unlock()
			return
		}

		printNumber(z.counter)
		z.counter += 1
		z.zero = true
		z.cond.Broadcast()
		z.cond.L.Unlock()
	}
}
