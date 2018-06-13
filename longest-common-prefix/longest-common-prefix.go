package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	res := ""

	i := 0

	if len(strs) > 0 {
		for true {
			var tmp byte
			for j := 0; j < len(strs); j++ {
				str := strs[j]

				if i < len(str) {
					b := str[i]
					if tmp == 0 {
						tmp = b
					} else if tmp != b {
						goto END
					}
				} else {
					goto END
				}
			}

			res += string(tmp)
			tmp = 0
			i++
		}
	}

END:
	return res
}
