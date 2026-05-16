package kth_largest_element

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthLargest(t *testing.T) {
	// func main() {
	// 	kthLargest := Constructor(3, []int{4, 5, 8, 2})
	// 	fmt.Println(kthLargest.Add(3))  // return 4
	// 	fmt.Println(kthLargest.Add(5))  // return 5
	// 	fmt.Println(kthLargest.Add(10)) // return 5
	// 	fmt.Println(kthLargest.Add(9))  // return 8
	// 	fmt.Println(kthLargest.Add(4))  // return 8

	// 	fmt.Println()
	// 	kthLargest = Constructor(4, []int{7, 7, 7, 7, 8, 3})
	// 	fmt.Println(kthLargest.Add(2))
	// 	fmt.Println(kthLargest.Add(10))
	// 	fmt.Println(kthLargest.Add(9))
	// 	fmt.Println(kthLargest.Add(9))
	// }

	tests := []struct {
		k        int
		starting []int
		ops      []int
		results  []int
	}{
		{
			k:        3,
			starting: []int{4, 5, 8, 2},
			ops:      []int{3, 5, 10, 9, 4},
			results:  []int{4, 5, 5, 8, 8},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %+v", tt.k, tt.starting), func(t *testing.T) {
			kthLargest := Constructor(3, []int{4, 5, 8, 2})
			results := []int{}
			for v := range tt.ops {
				results = append(results, kthLargest.Add(v))
			}

			if !assert.Equal(
				t,
				tt.results,
				results,
			) {
				t.Errorf("wanted %+v, got %+v", tt.results, results)
			}

			// testify
			// fmt.Println(kthLargest.Add(3))  // return 4
			// fmt.Println(kthLargest.Add(5))  // return 5
			// fmt.Println(kthLargest.Add(10)) // return 5
			// fmt.Println(kthLargest.Add(9))  // return 8
			// fmt.Println(kthLargest.Add(4))  // return 8
		})
	}
}
