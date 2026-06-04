package solution

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		slice_l1  []int
		slice_l2  []int
		slice_res []int
	}{
		{
			slice_l1:  []int{2, 4, 3},
			slice_l2:  []int{5, 6, 4},
			slice_res: []int{7, 0, 8},
		},
		{
			slice_l1:  []int{0},
			slice_l2:  []int{0},
			slice_res: []int{0},
		},
		{
			slice_l1:  []int{2, 4, 3},
			slice_l2:  []int{5, 6},
			slice_res: []int{7, 0, 4},
		},

		{
			slice_l1:  []int{9, 9, 9, 9},
			slice_l2:  []int{9, 9, 9, 9, 9, 9, 9},
			slice_res: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v + %+v", tt.slice_l1, tt.slice_l2), func(t *testing.T) {
			// t.Parallel()
			res := addTwoNumbers(
				sToN(t, tt.slice_l1),
				sToN(t, tt.slice_l2),
			)

			var ln_res *ListNode = res
			res_slice := []int{}

			for ln_res != nil {
				// if ln_res.Val != ln_want.Val {
				// 	t.Errorf("error value: wanted %d, got %d", ln_want.Val, ln_res.Val)
				// }
				res_slice = append(res_slice, ln_res.Val)
				ln_res = ln_res.Next
			}

			assert.Equal(t, res_slice, tt.slice_res)

			// if res.Next == nil && ln_want.Next != nil {
			// 	t.Errorf("different null states %+v %+v", res, ln_want)

			// }
		})
	}
}

func sToN(t *testing.T, n []int) *ListNode {
	t.Helper()
	ln := &ListNode{
		Val: n[0],
	}
	or_ln := ln
	if len(n) == 1 {
		return ln
	}
	for i, num := range n {
		if i == 0 {
			continue
		}
		ln2 := &ListNode{
			Val: num,
		}
		ln.Next = ln2
		ln = ln2
	}

	return or_ln
}
