package asteroidcolissions

import (
	"errors"
)

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

func isGoingRight(i int) bool {
	return i > 0
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}

func willCollide(s *Stack, a int) bool {
	prev, err := s.Peek()
	if err != nil {
		panic(err)
	}

	if (isGoingRight(a) && isGoingRight(prev)) ||
		(!isGoingRight(a) && !isGoingRight(prev)) ||
		!isGoingRight(prev) && isGoingRight(a) {

		return false
	}

	return true
}

func asteroidCollision(asteroids []int) []int {
	s := NewStack()

	for i, a := range asteroids {
		if i == 0 {
			s.Push(a)
			// log.Printf("pushing %d\n", a)
			continue
		}

		prev, err := s.Peek()
		if err != nil {
			s.Push(a)
			continue
		}
		if !willCollide(s, a) {
			// log.Printf("pushing %d\n", a)
			// log.Printf("%d and %d will not collide\n", prev, a)
			s.Push(a)
			continue
		} else {
			for willCollide(s, a) {
				// log.Printf("%d and %d will collide\n", prev, a)
				prev, err = s.Pop()
				if err != nil {
					panic(err)
				}
				if abs(prev) > abs(a) {
					// log.Printf("%d beats %d", prev, a)
					s.Push(prev)
					break
				} else if abs(prev) == abs(a) {
					// log.Printf("%d mutually-destroyed %d", prev, a)
					break
				} else {
					// log.Printf("%d destroyed %d", a, prev)
					if s.IsEmpty() {
						// log.Printf("stack empty pushing")
						s.Push(a)
						break
					}

					if !willCollide(s, a) {
						s.Push(a)
					}
				}
			}
		}
	}

	return s.s
}
