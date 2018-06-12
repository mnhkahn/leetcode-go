package reverse_integer

import "math"

func reverse(x int) int {
	res := 0
	for x != 0 {
		y := x / 10
		remainder := x - y*10

		res = res*10 + remainder

		x = y
	}

	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return res
}
