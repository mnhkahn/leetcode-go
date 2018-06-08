package two_sum

import "sort"

func twoSum(nums []int, target int) []int {
	temp := make([]int, len(nums), len(nums))
	copy(temp, nums)

	tempUseFlag := make([]bool, len(nums), len(nums))

	sort.Ints(temp)

	a, b := -1, -1

	for i := 0; i < len(temp)-1; i++ {
		for j := 1; j < len(temp); j++ {
			if temp[i]+temp[j] == target {
				a = temp[i]
				b = temp[j]
				goto END
			} else if temp[i]+temp[j] < target {
				continue
			} else {
				break
			}
		}
	}

END:
	if a != -1 && b != -1 {
		return []int{findIndex(nums, tempUseFlag, a), findIndex(nums, tempUseFlag, b)}
	}

	return nil
}

func findIndex(nums []int, tempUseFlag []bool, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val && !tempUseFlag[i] {
			tempUseFlag[i] = true
			return i
		}
	}
	return -1
}
