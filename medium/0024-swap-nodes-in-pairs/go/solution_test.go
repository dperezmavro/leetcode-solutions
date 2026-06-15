package swap_nodes_in_pairs

import (
	"fmt"
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := swapPairs(LNFromArr(t, tt.input))
			res_arr := ArrFromLN(t, res)
			assert.Equal(t, tt.want, res_arr)
		})
	}
}

func LNFromArr(t *testing.T, in []int) *ListNode {
	t.Helper()

	if in == nil {
		return nil
	}

	if len(in) == 1 {
		return &ListNode{
			Val: in[0],
		}
	}

	var h *ListNode
	head := h
	for _, v := range in {
		if h != nil {
			h.Next = &ListNode{
				Val: v,
			}
			h = h.Next
		} else {
			h = &ListNode{
				Val: v,
			}
		}
	}

	return head
}

func ArrFromLN(t *testing.T, h *ListNode) []int {
	t.Helper()

	if h == nil {
		return nil
	}

	fmt.Printf("returnning array \n")
	res := []int{}

	for h != nil {
		res = append(res, h.Val)
		h = h.Next
	}

	fmt.Printf("returnning array %+v\n", res)

	return res
}
