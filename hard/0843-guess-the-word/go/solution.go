package solution

import (
	"slices"
)

var stop bool = false
var checkedAlready map[string]bool = make(map[string]bool, 10)
var samples = 2

func findSecretWord(words []string, master *Master) {
	maxMatch := 0
	maxWord := ""
	for i := range samples {
		idx := i * (len(words) / samples)
		if _, ok := checkedAlready[words[idx]]; ok {
			idx = ((i + samples) * (len(words) / samples)) % len(words)
		}
		res := master.Guess(words[idx])
		checkedAlready[words[idx]] = true
		if res == 6 {
			stop = true
			return
		} else if res == -1 {
			words = slices.Delete(words, idx, idx+1)
		} else if res > maxMatch {
			maxWord = words[idx]
			maxMatch = res
			words = slices.Delete(words, idx, idx+1)
		}
	}

	words = reduceWordSet(maxWord, maxMatch, words)
	if !stop {
		findSecretWord(words, master)
	}
}

func reduceWordSet(maxWord string, maxMatch int, words []string) []string {
	d := map[string]int{}
	for _, w := range words {
		d[w] = manhattanDistance(maxWord, w)
	}

	for k, v := range d {
		if v > (6 - maxMatch) {
			delete(d, k)
		}
	}

	res := []string{}
	for w := range d {
		res = append(res, w)
	}
	return res
}

func manhattanDistance(s1, s2 string) int {
	distance := 0 // identical strings
	for i := range len(s1) {
		if s1[i] != s2[i] {
			distance++
		}
	}

	return distance
}
