package median_of_two_sorted_arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := merge(nums1, nums2)

	if len(nums) == 0 {
		return 0
	}

	if len(nums)%2 == 1 {
		return float64(nums[len(nums)/2])
	} else {
		return float64(nums[len(nums)/2]+nums[len(nums)/2-1]) / 2
	}
}

func merge(nums1, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	} else if len(nums2) == 0 {
		return nums1
	}

	mergeIds := make([]int, 0, len(nums1)+len(nums2))

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			mergeIds = append(mergeIds, nums2[j])
			j++
		} else {
			mergeIds = append(mergeIds, nums1[i])
			i++
		}
	}

	for ; i < len(nums1); i++ {
		mergeIds = append(mergeIds, nums1[i])
	}

	for ; j < len(nums2); j++ {
		mergeIds = append(mergeIds, nums2[j])
	}

	return mergeIds
}
