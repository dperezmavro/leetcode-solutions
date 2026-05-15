package asteroidcolissions

import "errors"

type Stack struct {
	s []int
}

func NewStack() *Stack {
	return &Stack{
		s: []int{},
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.s) == 0
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	r := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]

	return r, nil
}

func (s *Stack) Push(a int) {
	s.s = append(s.s, a)
}

func (s *Stack) Peek() (int, error) {
	if !s.IsEmpty() {
		return s.s[len(s.s)-1], nil
	}

	return 0, errors.New("stack is empty")
}
