package simplifypath

import "testing"

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{
			input:  "/home/",
			output: "/home",
		},
		{
			input:  "/home//foo/",
			output: "/home/foo",
		},
		{
			input:  "/home/user/Documents/../Pictures",
			output: "/home/user/Pictures",
		},
		{
			input:  "/../",
			output: "/",
		},
		{
			input:  "/.../a/../b/c/../d/./",
			output: "/.../b/d",
		},
		{
			input:  "/a/../../b/../c//.//",
			output: "/c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := simplifyPath(tt.input)
			if tt.output != result {
				t.Errorf("error: wanted %s, got %s", tt.output, result)
			}
		})
	}

}
