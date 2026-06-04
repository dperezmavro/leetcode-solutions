package solution

import (
	"log"
	"slices"
)

func findSecretWord(words []string, master *Master) {
	var checkedAlready map[string]bool = make(map[string]bool, 10)
	var distances map[string]int = make(map[string]int)
	p := &PasswordFinder{
		words:        words,
		master:       master,
		distances:    distances,
		checkedWords: checkedAlready,
	}
	p.findSecretWordReturns()
}

type PasswordFinder struct {
	words        []string
	checkedWords map[string]bool
	distances    map[string]int
	master       *Master
}

func (p *PasswordFinder) removeWordsWithDLT(
	d int,
	word string,
) {
	result := []string{}
	for _, w := range p.words {
		dd := p.manhattanDistance(word, w)
		if dd >= d {
			result = append(result, w)
		}
	}
	p.words = result
}

func (p *PasswordFinder) removeWordsWithDGT(
	d int,
	word string,
) {
	result := []string{}

	for _, w := range p.words {
		dd := p.manhattanDistance(word, w)
		if dd <= 6-d {
			result = append(result, w)
		}
	}
	p.words = result
}

func (p *PasswordFinder) manhattanDistance(s1, s2 string) int {
	distance := 0 // identical strings
	if d, ok := p.distances[s1+s2]; ok {
		return d
	}
	for i := range len(s1) {
		if s1[i] != s2[i] {
			distance++
		}
	}

	// same distance
	p.distances[s1+s2] = distance
	p.distances[s2+s1] = distance
	return distance
}

// convenience function
func (p *PasswordFinder) findSecretWordReturns() string {
	idx := len(p.words) / 2

	passwordDelta := p.master.Guess(p.words[idx])
	p.checkedWords[p.words[idx]] = true

	switch passwordDelta {
	case 6:
		return p.words[idx]
	case -1:
		word_to_remove := p.words[idx]
		p.words = slices.Delete(p.words, idx, idx+1)
		p.removeWordsWithDLT(6, word_to_remove)

		return p.findSecretWordReturns()
	default:
		currentWord := p.words[idx]

		// delete current word, not a solution
		p.words = slices.Delete(p.words, idx, idx+1)
		// remove words with a D less than the delta needed to reach the password
		// i.e. words too similar to this word, but too dissimilar to the password
		p.removeWordsWithDLT(6-passwordDelta, currentWord)

		// remove words with a delta
		p.removeWordsWithDGT(passwordDelta, currentWord)
		log.Printf("filter default-2: %d: %+v", len(p.words), p.words)

		return p.findSecretWordReturns()
	}
}
