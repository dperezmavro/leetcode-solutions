package swap_nodes_in_pairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoluion(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "empty",
			input: []int{},
			want:  nil,
		},
		{
			name:  "nil",
			input: nil,
			want:  nil,
		},
		{
			name:  "one",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "three",
			input: []int{1, 2, 3},
			want:  []int{2, 1, 3},
		},
		{
			name:  "four",
			input: []int{1, 2, 3, 4},
			want:  []int{2, 1, 4, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// to test points
			res := swapPairs(LNFromArr(t, tt.input))
			res_arr := ArrFromLN(t, res)
			assert.Equal(t, tt.want, res_arr)

			// to test values
			res_2 := swapPairsValues(LNFromArr(t, tt.input))
			res_arr_2 := ArrFromLN(t, res_2)
			assert.Equal(t, tt.want, res_arr_2)
		})
	}
}

func LNFromArr(t *testing.T, in []int) *ListNode {
	t.Helper()

	if len(in) == 0 {
		return nil
	}

	head := &ListNode{Val: in[0]}
	curr := head
	for _, v := range in[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}

	return head
}

func ArrFromLN(t *testing.T, h *ListNode) []int {
	t.Helper()

	if h == nil {
		return nil
	}

	res := []int{}
	for h != nil {
		res = append(res, h.Val)
		h = h.Next
	}

	return res
}
