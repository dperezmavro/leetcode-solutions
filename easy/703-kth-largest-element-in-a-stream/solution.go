package kth_largest_element

import (
	"container/heap"
)

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

func (k *KthLargest) Add(val int) int {
	heap.Push(k.minHeap, val)

	if k.minHeap.Len() > k.k {
		heap.Pop(k.minHeap)
	}
	return (*k.minHeap)[0]
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
