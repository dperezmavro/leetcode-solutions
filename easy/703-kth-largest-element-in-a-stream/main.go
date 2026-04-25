package main

import (
	"container/heap"
	"fmt"
)

func main() {
	kthLargest := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kthLargest.Add(3))  // return 4
	fmt.Println(kthLargest.Add(5))  // return 5
	fmt.Println(kthLargest.Add(10)) // return 5
	fmt.Println(kthLargest.Add(9))  // return 8
	fmt.Println(kthLargest.Add(4))  // return 8

	fmt.Println()
	kthLargest = Constructor(4, []int{7, 7, 7, 7, 8, 3})
	fmt.Println(kthLargest.Add(2))
	fmt.Println(kthLargest.Add(10))
	fmt.Println(kthLargest.Add(9))
	fmt.Println(kthLargest.Add(9))
}

type KthLargest struct {
	k       int
	nums    []int
	minHeap *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	kth := KthLargest{
		k:       k,
		nums:    nums,
		minHeap: minHeap,
	}
	for _, n := range nums {
		kth.Add(n)
	}
	return kth
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.minHeap, val)

	if this.minHeap.Len() > this.k {
		heap.Pop(this.minHeap)
	}
	return (*this.minHeap)[0]
}

type MinHeap []int

func (p MinHeap) Len() int           { return len(p) }
func (p MinHeap) Less(i, j int) bool { return p[i] < p[j] }
func (p MinHeap) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *MinHeap) Push(n any) {
	*p = append(*p, n.(int))
}

func (p *MinHeap) Pop() any {
	old := *p
	n := len(old)
	num := old[n-1]
	*p = old[:n-1]
	return num
}
