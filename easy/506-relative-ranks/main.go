package main

import (
	"container/heap"
	"fmt"
	"log"
)

func main() {
	log.Printf("%+v\n", findRelativeRanks([]int{5, 4, 3, 2, 1}))
	log.Printf("%+v\n", findRelativeRanks([]int{10, 3, 8, 9, 4}))
}

type AthleteQ []int

func (a *AthleteQ) Pop() any {
	old := *a
	n := len(old)
	res := old[n-1]
	*a = old[0 : n-1]

	return res
}
func (a AthleteQ) Len() int { return len(a) }

// this is for a max heap
func (a AthleteQ) Less(x, y int) bool { return a[x] > a[y] }
func (a *AthleteQ) Push(x any)        { *a = append(*a, x.(int)) }
func (a AthleteQ) Swap(x, y int)      { a[x], a[y] = a[y], a[x] }

func findRelativeRanks(score []int) []string {
	answer := make([]string, len(score))

	q := &AthleteQ{}

	heap.Init(q)
	for _, n := range score {
		heap.Push(q, n)
	}

	place := 1
	rankings := map[int]string{}
	for q.Len() > 0 {
		score := heap.Pop(q).(int)
		log.Printf("score %d\n", score)
		if place == 1 {
			rankings[score] = "Gold Medal"
		} else if place == 2 {
			rankings[score] = "Silver Medal"
		} else if place == 3 {
			rankings[score] = "Bronze Medal"
		} else {
			rankings[score] = fmt.Sprintf("%d", place)
		}
		place++
	}

	for i, v := range score {
		answer[i] = rankings[v]
	}

	return answer
}
