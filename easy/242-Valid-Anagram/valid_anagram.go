package valid_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sCount := getCount(s)
	tCount := getCount(t)

	for k, v := range sCount {
		if tCount[k] != v {
			return false
		}
	}

	return true
}

func getCount(s string) map[rune]int {
	sCounts := map[rune]int{}
	for _, c := range s {
		sCounts[c]++
	}

	return sCounts
}
