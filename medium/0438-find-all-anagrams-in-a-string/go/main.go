package find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	var res []int = []int{}
	if len(s) < len(p) {
		return res
	}

	pCounts := getCount(p)

	windowCounts := getCount(s[0:len(p)])
	if checkMaps(windowCounts, pCounts) {
		res = append(res, 0)
	}
	var i int = 0
	for i = 1; i < len(s)-len(p)+1; i++ {
		r := rune(s[i-1])
		r_now := rune(s[i+len(p)-1])
		windowCounts[r]--
		windowCounts[r_now]++
		if checkMaps(windowCounts, pCounts) {
			res = append(res, i)
		}
	}

	return res
}

func checkMaps(s, p map[rune]int) bool {
	for k, v := range p {
		if s[k] != v {
			return false
		}
	}

	return true
}

func getCount(s string) map[rune]int {
	sCounts := make(map[rune]int, len(s))
	for _, c := range s {
		sCounts[c]++
	}

	return sCounts
}
