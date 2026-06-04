package solution

import (
	"log"
	"slices"
)

func findSecretWord(words []string, master *Master) {
	var checkedAlready map[string]bool = make(map[string]bool, 10)
	findSecretWordReturns(words, master, checkedAlready)
}
func findSecretWordReturns(words []string, master *Master, checkedAlready map[string]bool) string {
	idx := len(words) / 2
	for _, ok := checkedAlready[words[idx]]; ok; {
		// in case we have checked this word, pick another word
		idx = (len(words) + 1) / 2
	}

	res := master.Guess(words[idx])
	// log.Printf("Guessed word: %s: %d\n", words[idx], res)

	checkedAlready[words[idx]] = true

	switch res {
	case 6:
		return words[idx]
	case -1:
		// word_to_remove := words[idx]
		words = slices.Delete(words, idx, idx+1)
		return findSecretWordReturns(words, master, checkedAlready)
	}

	currentWord := words[idx]
	currentWordMatch := res

	// delete current word, not a solution
	words = slices.Delete(words, idx, idx+1)

	// delete all words whose delta is less than 6 - match
	// iterate backwards so we don't delete the wrong word
	// for i := len(words) - 1; i >= 0; i-- {
	// 	d := manhattanDistance(currentWord, words[i])
	// 	if d < 6-currentWordMatch {
	// 		log.Printf("deleting word %s %d\n", words[i], d)
	// 		words = slices.Delete(words, i, i+1)
	// 	}
	// }
	words = removeWordsWithDLT(6-currentWordMatch, words, currentWord)

	words = reduceWordSet(currentWord, currentWordMatch, words)
	return findSecretWordReturns(words, master, checkedAlready)
}

func removeWordsWithDLT(d int, words []string, word string) []string {
	result := []string{}
	for _, w := range words {
		dd := manhattanDistance(word, w)
		if dd >= d {
			result = append(result, w)
			log.Printf("keeping word %s %d\n", w, dd)
		} else {
			log.Printf("removing word %s %d\n", w, dd)
		}
	}

	log.Printf("original words: %d, new words %d\n", len(words), len(result))

	return result
}

func reduceWordSet(currentWord string, currentMax int, words []string) []string {
	d := map[string]int{}
	for _, w := range words {
		if w == currentWord {
			d[w] = currentMax
			continue
		}
		d[w] = manhattanDistance(currentWord, w)
	}

	for k, v := range d {
		if v > (6 - currentMax) {
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
