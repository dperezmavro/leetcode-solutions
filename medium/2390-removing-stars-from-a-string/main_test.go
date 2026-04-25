package removingstarsfromastring

import "testing"

func TestRemoveStars(t *testing.T) {
	tests := []struct {
		want  string
		input string
	}{
		{
			want:  "lecoe",
			input: "leet**cod*e",
		},
		{
			want:  "",
			input: "erase*****",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			res := removeStars(tt.input)
			if res != tt.want {
				t.Errorf("wanted %s, got %s", tt.want, res)
			}
		})
	}
}
