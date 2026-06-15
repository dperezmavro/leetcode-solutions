package min_stack

import "container/heap"

type MinStack struct {
	stack *IntHeap
}

func Constructor() MinStack {
	s := &IntHeap{}
	heap.Init(s)
	return MinStack{
		stack: s,
	}
}

func (m *MinStack) Push(value int) {
	heap.Push(m.stack, value)
}

func (m *MinStack) Pop() {
	heap.Pop(m.stack)
}

func (m *MinStack) Top() int {
	if m.stack.Len() > 0 {
		return (*m.stack)[m.stack.Len()-1]
	}
	return -1
}

func (m *MinStack) GetMin() int {
	if m.stack.Len() > 0 {
		return (*m.stack)[0]
	}
	return -1
}

type IntHeap []int

func (i IntHeap) Len() int           { return len(i) }
func (i IntHeap) Swap(x, y int)      { i[x], i[y] = i[y], i[x] }
func (i IntHeap) Less(x, y int) bool { return i[x] < i[y] }
func (i *IntHeap) Pop() any {
	old := *i
	x := old[len(old)-1]
	n := len(old)
	*i = old[0 : n-1]
	return x
}

func (i *IntHeap) Push(x any) {
	*i = append(*i, x.(int))
}
