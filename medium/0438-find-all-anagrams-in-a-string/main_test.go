package find_all_anagrams_in_a_string

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		s   string
		p   string
		out []int
	}{
		{
			s:   "cbaebabacd",
			p:   "abc",
			out: []int{0, 6},
		},
		{
			s:   "cbdaebabacd",
			p:   "abc",
			out: []int{7},
		},
		{
			s:   "abab",
			p:   "ab",
			out: []int{0, 1, 2},
		},
		{
			s:   "ababab",
			p:   "abaaaaaa",
			out: []int{},
		},
		{
			s:   "baa",
			p:   "aa",
			out: []int{1},
		},
		{
			s:   "acdcaeccde",
			p:   "c",
			out: []int{1, 3, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			out := findAnagrams(tt.s, tt.p)
			if !reflect.DeepEqual(out, tt.out) {
				t.Errorf("failed for %s: wanted %+v, got %+v", tt.s, tt.out, out)
			}
		})
	}
}
