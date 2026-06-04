package solution

func findSecretWord(words []string, master *Master) {
	minGuesses := 10
	if len(words) <= minGuesses {
		for _, w := range words {
			master.Guess(w)
		}

		return
	}

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
