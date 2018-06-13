package _sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test3Sum(t *testing.T) {
	assert.Equal(t, [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
	assert.Equal(t, nil, threeSum([]int{0, 1, 1}))
	assert.Equal(t, [][]int{[]int{-2, 0, 2}, []int{-2, 1, 1}}, threeSum([]int{-2, 0, 1, 1, 2}))
}
