package divide_two_integers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivideTwoInts(t *testing.T) {
	assert.Equal(t, 3, divide(10, 3))
	assert.Equal(t, -3, divide(7, -2))
	assert.Equal(t, 0, divide(0, 1))
	assert.Equal(t, 1, divide(1, 1))
	assert.Equal(t, 2147483647, divide(-2147483648, -1))
	assert.Equal(t, -2147483648, divide(-2147483648, 1))
	assert.Equal(t, 1073741823, divide(2147483647, 2))
}
