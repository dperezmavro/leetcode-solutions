package solution

import (
	"slices"
)

func findSecretWord(words []string, master *Master) {
	var checkedAlready map[string]bool = make(map[string]bool, 10)
	var distances map[string]int = make(map[string]int)
	findSecretWordReturns(words, master, checkedAlready, distances)
}

// convenience function
func findSecretWordReturns(
	words []string,
	master *Master,
	checkedAlready map[string]bool,
	distances map[string]int,
) string {
	idx := len(words) / 2
	for _, ok := checkedAlready[words[idx]]; ok; {
		// in case we have checked this word, pick another word
		idx = (len(words) + 1) / 2
	}

	passwordDelta := master.Guess(words[idx])
	checkedAlready[words[idx]] = true

	switch passwordDelta {
	case 6:
		return words[idx]
	case -1:
		word_to_remove := words[idx]
		words = slices.Delete(words, idx, idx+1)
		words = removeWordsWithDLT(6, words, word_to_remove, distances)
		// log.Printf("filter -1: %d: %+v", len(words), words)
		return findSecretWordReturns(words, master, checkedAlready, distances)
	default:
		currentWord := words[idx]

		// delete current word, not a solution
		words = slices.Delete(words, idx, idx+1)
		// remove words with a D less than the delta needed to reach the password
		// i.e. words too similar to this word, but too dissimilar to the password
		words = removeWordsWithDLT(6-passwordDelta, words, currentWord, distances)

		// log.Printf("filter default-1: %d: %+v", len(words), words)

		// remove words with a delta
		words = removeWordsWithDGT(passwordDelta, words, currentWord, distances)
		// log.Printf("filter default-2: %d: %+v", len(words), words)

		return findSecretWordReturns(words, master, checkedAlready, distances)
	}
}

// removeWordsWithDLT removes words whose manhattan
// distance is less than d
func removeWordsWithDLT(
	d int,
	words []string,
	word string,
	distances map[string]int,
) []string {
	result := []string{}
	for _, w := range words {
		dd := manhattanDistance(word, w, distances)
		if dd >= d {
			result = append(result, w)
		}
	}
	return result
}

// removeWordsWithDGT removes words whose manhattan
// distance is greater than 6-d
func removeWordsWithDGT(
	d int,
	words []string,
	word string,
	distances map[string]int,
) []string {
	result := []string{}

	for _, w := range words {
		dd := manhattanDistance(word, w, distances)
		if dd <= 6-d {
			result = append(result, w)
		}
	}
	return result
}

func manhattanDistance(s1, s2 string, distances map[string]int) int {
	distance := 0 // identical strings
	if d, ok := distances[s1+s2]; ok {
		return d
	}
	for i := range len(s1) {
		if s1[i] != s2[i] {
			distance++
		}
	}

	distances[s1+s2] = distance

	return distance
}
