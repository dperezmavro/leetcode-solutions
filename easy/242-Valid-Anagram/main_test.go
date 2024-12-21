package valid_anagram

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		s   string
		t   string
		res bool
	}{
		{
			s:   "anagram",
			t:   "nagaram",
			res: true,
		},
		{
			s: "car",
			t: "rat",
		},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if isAnagram(tt.s, tt.t) != tt.res {
				t.Errorf("failed for %+v", tt)
			}
		})
	}
}
