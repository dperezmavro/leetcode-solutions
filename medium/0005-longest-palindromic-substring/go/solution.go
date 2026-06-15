package logest_palindromic_substring

func longestPalindrome(s string) string {
	max_l := 0
	max_r := 0

	for i := 0; i < len(s); i++ {
		odd_l, odd_r := expand_centre(s, i, i)
		even_l, even_r := expand_centre(s, i, i+1)
		if odd_r-odd_l > max_r-max_l {
			max_l, max_r = odd_l, odd_r
		}

		if even_r-even_l > max_r-max_l {
			max_l, max_r = even_l, even_r
		}
	}

	return s[max_l : max_r+1]
}

func expand_centre(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left -= 1
		right += 1
	}
	return left + 1, right - 1
}
