package asteroidcolissions

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
			continue
		}

		_, err := s.Peek()
		if err != nil {
			s.Push(a)
			continue
		}
		if !willCollide(s, a) {
			s.Push(a)
			continue
		} else {
			for willCollide(s, a) {
				prev, err := s.Pop()
				if err != nil {
					panic(err)
				}
				if abs(prev) > abs(a) {
					s.Push(prev)
					break
				} else if abs(prev) == abs(a) {
					break
				} else {
					if s.IsEmpty() {
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
