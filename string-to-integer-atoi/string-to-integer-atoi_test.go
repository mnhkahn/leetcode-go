package string_to_integer_atoi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {
	assert.Equal(t, 42, myAtoi("42"))
	assert.Equal(t, -42, myAtoi("   -42"))
	assert.Equal(t, 4193, myAtoi("4193 with words"))
	assert.Equal(t, 0, myAtoi("words and 987"))
	assert.Equal(t, -2147483648, myAtoi("-91283472332"))
	assert.Equal(t, 3, myAtoi("3.14159"))
	assert.Equal(t, -3, myAtoi("-3.14159"))
	assert.Equal(t, 3, myAtoi("3+3"))
	assert.Equal(t, 0, myAtoi("+-2"))
	assert.Equal(t, 2147483647, myAtoi("9223372036854775808"))
	assert.Equal(t, 0, myAtoi("-   234"))
}
