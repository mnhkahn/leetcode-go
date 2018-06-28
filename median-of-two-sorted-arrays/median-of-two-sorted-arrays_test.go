package median_of_two_sorted_arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedianOfTwoSortedArrays(t *testing.T) {
	assert.Equal(t, 2.0, findMedianSortedArrays([]int{1, 3}, []int{2}))
	assert.Equal(t, 2.5, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, 3.5, findMedianSortedArrays(nil, []int{3, 4}))
}
