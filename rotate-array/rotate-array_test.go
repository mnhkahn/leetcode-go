package rotate_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateArray(t *testing.T) {
	//a := []int{1, 2, 3, 4, 5, 6, 7}
	//rotate(a, 1)
	//assert.Equal(t, []int{7, 1, 2, 3, 4, 5, 6}, a)
	//
	//b := []int{1, 2, 3, 4, 5, 6, 7}
	//rotate(b, 3)
	//assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, b)

	c := []int{-1, -100, 3, 99}
	rotate(c, 2)
	assert.Equal(t, []int{3, 99, -1, -100}, c)
}
