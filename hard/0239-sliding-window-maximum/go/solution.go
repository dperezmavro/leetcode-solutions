package sliding_window_maximum

import (
	"container/heap"
)

type MaxHeap []int

func (m MaxHeap) Len() int           { return len(m) }
func (m MaxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MaxHeap) Less(i, j int) bool { return m[i] > m[j] }
func (m *MaxHeap) Push(i any)        { *m = append(*m, i.(int)) }
func (m *MaxHeap) Pop() any {
	old := *m
	x := old[len(old)-1]
	old = old[:len(old)-1]
	*m = old
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	// special base case
	if k == 1 {
		return nums
	}

	m := make(MaxHeap, 0)
	heap.Init(&m)

	result := []int{}
	currMax := 0

	for _, c := range nums[0:k] {
		heap.Push(&m, c)
	}
	currWindow := nums[0:k]
	currMax = m[0]

	result = append(result, currMax)
	for _, n := range nums[k:] {
		oldestNum := currWindow[0]
		// if the top of the heap and oldest number and biggest number, remove
		if m[0] == oldestNum {
			heap.Pop(&m)
		} else {
			// need to find where to pop the number from
			idx := 0
			for i, v := range m {
				if v == oldestNum {
					idx = i
					break
				}
			}
			heap.Remove(&m, idx)
		}
		currWindow = currWindow[1:]
		currWindow = append(currWindow, n)

		heap.Push(&m, n)
		currMax = m[0]

		result = append(result, currMax)
	}
	return result
}
