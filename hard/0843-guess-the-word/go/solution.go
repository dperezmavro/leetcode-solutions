package solution

func findSecretWord(words []string, master *Master) {
	p := &PasswordFinder{
		words:     words,
		master:    master,
		distances: make(map[string]int),
	}
	p.findSecretWordReturns()
}

type PasswordFinder struct {
	words     []string
	distances map[string]int
	master    *Master
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

// bestWord picks the candidate that minimises the worst-case remaining
// candidates over all possible match counts (minimax).
func (p *PasswordFinder) bestWord() string {
	best := p.words[0]
	bestWorst := len(p.words) + 1
	for _, candidate := range p.words {
		var buckets [7]int
		for _, w := range p.words {
			if w == candidate {
				continue
			}
			buckets[p.manhattanDistance(candidate, w)]++
		}
		worst := 0
		for _, count := range buckets {
			if count > worst {
				worst = count
			}
		}
		if worst < bestWorst {
			bestWorst = worst
			best = candidate
		}
	}

	return best
}

// convenience function
func (p *PasswordFinder) findSecretWordReturns() string {
	currentWord := p.bestWord()
	passwordDelta := p.master.Guess(currentWord)

	switch passwordDelta {
	case 6:
		return currentWord
	case -1:
		p.removeWordsWithDLT(6, currentWord)

		return p.findSecretWordReturns()
	default:
		// remove words with a D less than the delta needed to reach the password
		// i.e. words too similar to this word, but too dissimilar to the password
		p.removeWordsWithDLT(6-passwordDelta, currentWord)

		// remove words with a D greater than the difference needed to reach the password
		// i.e. words too dissimilar to this word
		p.removeWordsWithDGT(passwordDelta, currentWord)

		return p.findSecretWordReturns()
	}
}
