package main

import (
	"fmt"
	"log"
	"sync"
)

type FooBar struct {
	n   int
	foo chan bool
	bar chan bool
}

func NewFooBar(n int) *FooBar {
	barChan := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Go(func() {
		barChan <- true
	})

	return &FooBar{
		n:   n,
		foo: make(chan bool),
		bar: barChan,
	}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.bar
		printFoo()
		fb.foo <- true
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.foo
		printBar()
		if i == fb.n-1 {
			return
		}
		fb.bar <- true
	}
}

func main() {
	fb := NewFooBar(2)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		fb.Foo(func() {
			fmt.Printf("foo")
		})
		log.Println("Done1")

		wg.Done()
	}()
	go func() {
		fb.Bar(func() {
			fmt.Printf("bar\n")
		})
		log.Println("Done2")

		wg.Done()
	}()

	wg.Wait()
}
