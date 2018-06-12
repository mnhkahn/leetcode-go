package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	ss := []byte(s)

	longest := 1
	for i := 0; i < len(ss); i++ {
		temp := make(map[byte]struct{}, len(s))
		temp[ss[i]] = struct{}{}

		for j := i + 1; j < len(ss); j++ {
			sj := ss[j]
			if _, exists := temp[sj]; !exists {
				tempLongest := j - i + 1
				if tempLongest > longest {
					longest = tempLongest
				}
				temp[sj] = struct{}{}
			} else {
				break
			}
		}
	}

	return longest
}
