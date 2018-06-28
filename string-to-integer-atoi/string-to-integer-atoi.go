package string_to_integer_atoi

import "math"

func myAtoi(str string) int {
	gap := byte('0')

	res := 0
	strBytes := []byte(str)
	positiveFlag := 1
	hasPositiveFalg := false
	atoiIng := false

	for i := 0; i < len(str); i++ {
		b := strBytes[i]

		if b >= byte('0') && b <= byte('9') {
			a := strBytes[i] - gap
			newRes := res*10 + int(a)

			if positiveFlag == 1 && newRes > math.MaxInt32 {
				return math.MaxInt32
			} else if positiveFlag == -1 && newRes > math.MaxInt32 {
				return math.MinInt32
			}

			res = newRes

			atoiIng = true
		} else {
			if !atoiIng {
				if b == '-' {
					if hasPositiveFalg {
						return 0
					}
					positiveFlag = -1
					hasPositiveFalg = true
				} else if b == '+' {
					if hasPositiveFalg {
						return 0
					}
					positiveFlag = 1
					hasPositiveFalg = true
				} else if b == ' ' {
					if hasPositiveFalg {
						return 0
					}
				} else {
					return 0
				}
			} else { // atoiIng == true
				break
			}
		}
	}

	res = positiveFlag * res

	return res
}
