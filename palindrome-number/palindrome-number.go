package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	temps := []int{}
	y := x

	for y > 0 {
		z := y / 10
		temps = append(temps, y-10*z)

		y = z
	}

	i, j := 0, len(temps)-1

	for i < j {
		if temps[i] != temps[j] {
			return false
		}
		i++
		j--
	}

	return true
}
