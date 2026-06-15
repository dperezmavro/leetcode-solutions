package min_stack

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		assertFunc func(t *testing.T)
		name       string
	}{
		{
			name: "base",
			assertFunc: func(t *testing.T) {
				t.Helper()

				minStack := Constructor()
				minStack.Push(-2)
				minStack.Push(0)
				minStack.Push(-3)
				m := minStack.GetMin()
				if m != -3 {
					t.Errorf("wanted -3 for getMin, got %d", m)
				}
				minStack.Pop()
				m = minStack.Top()
				if m != 0 {
					t.Errorf("wanted 0 for top, got %d", m)
				}
				m = minStack.GetMin()
				if m != -2 {
					t.Errorf("wanted -2 for getMin, got %d", m)
				}
			},
		},
		{
			name: "test_1",
			assertFunc: func(t *testing.T) {
				t.Helper()

				minStack := Constructor()
				minStack.Push(-2)
				minStack.Push(0)
				minStack.Push(-1)
				m := minStack.GetMin()
				if m != -2 {
					t.Errorf("wanted -2 for getMin, got %d", m)
				}
				minStack.Pop()
				m = minStack.Top()
				if m != -1 {
					t.Errorf("wanted -1 for top, got %d", m)
				}
				m = minStack.GetMin()
				if m != -2 {
					t.Errorf("wanted -2 for getMin, got %d", m)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertFunc(t)
		})
	}
}
