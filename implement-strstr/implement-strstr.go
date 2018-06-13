package implement_strstr

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		tmp := haystack[i : i+len(needle)]
		if tmp == needle {
			return i
		}
	}

	return -1
}
