package linked_list_cycle

import (
	"testing"

	common "github.com/dperezmavro/leetcode-solutions/common/go"
	"github.com/magiconair/properties/assert"
)

func TestHasCycle(t *testing.T) {
	tests := []struct {
		getLL    func() *common.ListNode
		name     string
		hasCycle bool
	}{
		{
			getLL: func() *common.ListNode {
				return &common.ListNode{
					Val: -1,
				}
			},
			hasCycle: false,
			name:     "test1",
		},
		{
			hasCycle: true,
			getLL: func() *common.ListNode {
				ll1 := &common.ListNode{
					Val: 3,
				}
				ll2 := &common.ListNode{
					Val: 2,
				}
				ll3 := &common.ListNode{
					Val: 0,
				}
				ll4 := &common.ListNode{
					Val: -4,
				}

				ll1.Next = ll2
				ll2.Next = ll3
				ll3.Next = ll4
				ll4.Next = ll2

				return ll1
			},
			name: "test2",
		},
		{
			hasCycle: true,
			name:     "test3",
			getLL: func() *common.ListNode {
				ll1 := &common.ListNode{
					Val: 1,
				}
				ll2 := &common.ListNode{
					Val: 2,
				}

				ll1.Next = ll2
				ll2.Next = ll1

				return ll1
			},
		},

		{
			hasCycle: false,
			name:     "test4",
			getLL: func() *common.ListNode {
				ll1 := &common.ListNode{
					Val: 1,
				}
				ll2 := &common.ListNode{
					Val: 2,
				}

				ll1.Next = ll2

				return ll1
			},
		},

		{
			hasCycle: false,
			name:     "test5",
			getLL: func() *common.ListNode {
				ll1 := &common.ListNode{
					Val: 1,
				}
				ll2 := &common.ListNode{
					Val: 2,
				}
				ll3 := &common.ListNode{
					Val: 2,
				}
				ll4 := &common.ListNode{
					Val: 2,
				}

				ll1.Next = ll2
				ll2.Next = ll3
				ll3.Next = ll4

				return ll1
			},
		},

		{
			hasCycle: true,
			name:     "test6",
			getLL: func() *common.ListNode {
				ll1 := &common.ListNode{
					Val: 1,
				}
				ll2 := &common.ListNode{
					Val: 2,
				}
				ll3 := &common.ListNode{
					Val: 2,
				}
				ll4 := &common.ListNode{
					Val: 2,
				}

				ll1.Next = ll2
				ll2.Next = ll3
				ll3.Next = ll4
				ll4.Next = ll3

				return ll1
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasCycle := hasCycle(tt.getLL())
			assert.Equal(t, hasCycle, tt.hasCycle)
		})
	}
}
