package solution

import (
	"container/heap"
)

type Frequency struct {
	number int
	count  int
}

func topKFrequent(nums []int, k int) []int {
	counts := map[int]int{}
	for _, n := range nums {
		counts[n]++
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	for num, freq := range counts {
		heap.Push(pq, Frequency{
			number: num,
			count:  freq,
		})
	}

	res := []int{}
	for range k {
		el := heap.Pop(pq)
		el_f := el.(Frequency)
		res = append(res, el_f.number)
	}
	return res
}

type PriorityQueue []Frequency

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].count > pq[j].count } // use > for max-heap
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(i any) {
	f := i.(Frequency)
	(*pq) = append((*pq), f)
}

func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	old := *pq
	x := old[n-1]
	*pq = old[:n-1]

	return x
}
