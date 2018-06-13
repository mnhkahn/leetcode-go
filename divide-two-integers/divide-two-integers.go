package divide_two_integers

// http://www.voidcn.com/article/p-qyazcsex-hz.html

import "math"

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		panic("divisor can't be 0.")
	}

	if dividend > math.MaxInt32 {
		return math.MaxInt32
	} else if dividend < math.MinInt32 {
		return math.MinInt32
	}

	positiveFlag := true
	if dividend^divisor < 0 {
		positiveFlag = false
	}

	if divisor < 0 {
		divisor = opposite(divisor)
	}
	if dividend < 0 {
		dividend = opposite(dividend)
	}

	res := int(math.Floor(math.Exp(math.Log(float64(dividend)) - math.Log(float64(divisor)))))

	if res > math.MaxInt32 {
		if positiveFlag {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}

	if !positiveFlag {
		return opposite(res)
	}
	return res
}

func opposite(a int) int {
	return ^a + 1
}
