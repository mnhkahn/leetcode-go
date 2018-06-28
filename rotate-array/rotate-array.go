package rotate_array

func rotate(nums []int, k int) {
	if len(nums) == 0 {
		return
	} else if k%len(nums) == 0 {
		return
	}

	value := nums[0]
	j := 0
	for i := 0; i < len(nums); i++ {
		newPos := loc(nums, j, k)
		tmp := nums[newPos]
		nums[newPos] = value
		value = tmp
		j = newPos
	}
}

func loc(nums []int, i, k int) int {
	return (i + k + len(nums)) % len(nums)
}
