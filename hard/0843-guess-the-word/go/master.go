package solution

import (
	"log"
	"slices"
)

type Master struct {
	// for testing purposes
	k           int
	maxAttempts int
	pass        string
	words       []string
}

func (m *Master) Guess(word string) int {

	if m.k > m.maxAttempts {
		log.Println("Either you took too many guesses, or you did not find the secret word.")
		panic("too many attempts")
	}

	m.k++

	if !slices.Contains(m.words, word) {
		return -1
	} else {
		if word == m.pass {
			log.Println("You guessed the secret word correctly.")
			return len(m.pass)
		}
		matches := 0
		for i := range m.pass {
			if m.pass[i] == word[i] {
				matches += 1
			}
		}

		return matches
	}
}
