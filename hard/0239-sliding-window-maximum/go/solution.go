package sliding_window_maximum

import (
	"container/list"
)

func maxSlidingWindow(nums []int, k int) []int {
	// special base case
	if k == 1 {
		return nums
	}

	result := []int{}

	l := list.New()
	// build the first window
	for _, c := range nums[0:k] {
		if l.Len() == 0 {
			l.PushFront(c)
			continue
		}

		for l.Len() > 0 && l.Back().Value.(int) < c {
			l.Remove(l.Back())
		}

		l.PushBack(c)
	}
	currWindow := nums[0:k]

	result = append(result, l.Front().Value.(int))

	// iterate on the rest of the array
	for _, n := range nums[k:] {
		oldestNum := currWindow[0]
		if l.Front().Value.(int) == oldestNum {
			l.Remove(l.Front())
		}
		currWindow = currWindow[1:]
		currWindow = append(currWindow, n)

		for l.Len() > 0 && l.Back().Value.(int) < n {
			l.Remove(l.Back())
		}

		l.PushBack(n)
		result = append(result, l.Front().Value.(int))
	}
	return result
}
