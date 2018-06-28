package longest_palindromic_substring

func longestPalindrome(s string) string {
	for l := len(s); l > 0; l-- {
		for i := 0; i < len(s)-l+1; i++ {
			sub := s[i : i+l]
			if isPalindrome(sub) {
				return sub
			}
		}
	}
	return ""
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	i, j := 0, len(s)-1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}

	return true
}
