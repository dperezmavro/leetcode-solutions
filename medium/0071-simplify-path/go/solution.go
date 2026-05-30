package simplifypath

import (
	"strings"
)

func simplifyPath(path string) string {
	cleanPath := strings.ReplaceAll(path, "//", "/")
	cleanPath = strings.TrimRight(cleanPath, "/")

	parts := strings.Split(cleanPath, "/")

	finalPaths := []string{}

	for _, d := range parts {
		if d == "." || d == "" {
			continue
		} else if d == ".." {
			if len(finalPaths) > 0 {
				finalPaths = finalPaths[:len(finalPaths)-1]
			}
		} else {
			finalPaths = append(finalPaths, d)
		}
	}

	if len(finalPaths) == 0 {
		return "/"
	}

	return "/" + strings.Join(finalPaths, "/")
}
