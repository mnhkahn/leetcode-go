package _sum

// https://blog.csdn.net/q1242027878/article/details/54983416

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	var res [][]int

	tmp := make([]int, len(nums))
	copy(tmp, nums)

	sort.Ints(tmp)

	for i := 1; i < len(nums)-1; {
		m := 0
		n := len(nums) - 1
		for m < n && m < i && n > i {
			a := tmp[m]
			b := tmp[n]
			c := tmp[i]

			sum := a + b + c
			if sum == 0 {
				isExists := false
				for _, r := range res {
					if r[0] == a && r[1] == c && r[2] == b {
						isExists = true
						break
					}
				}
				if !isExists {
					res = append(res, []int{a, c, b})
				}
				m++
				n--
			} else {
				if a*b > 0 { // ab 同符号
					break
				} else if sum > 0 {
					n--
				} else if sum < 0 {
					m++
				}
			}
		}
		i++
	}

	return res

}
