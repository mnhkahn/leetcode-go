package roman_to_integer

var romanIntMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var speicalCaseMap = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func romanToInt(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {
		b := s[i]

		if i+1 < len(s) {
			if value, exists := speicalCaseMap[s[i:i+2]]; exists {
				res += value
				i++
				continue
			}
		}

		res = res + romanIntMap[b]
	}

	return res
}
