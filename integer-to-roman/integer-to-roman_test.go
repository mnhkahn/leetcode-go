package integer_to_roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	assert.Equal(t, "I", intToRoman(1))
	assert.Equal(t, "III", intToRoman(3))
	assert.Equal(t, "IV", intToRoman(4))
	assert.Equal(t, "V", intToRoman(5))
	assert.Equal(t, "XII", intToRoman(12))
	assert.Equal(t, "XXVII", intToRoman(27))
	assert.Equal(t, "IX", intToRoman(9))
	assert.Equal(t, "LVIII", intToRoman(58))
	assert.Equal(t, "XCIX", intToRoman(99))
	assert.Equal(t, "CMXCIV", intToRoman(994))
	assert.Equal(t, "MCMXCIV", intToRoman(1994))
	assert.Equal(t, "MMMCMXCIX", intToRoman(3999))
}
