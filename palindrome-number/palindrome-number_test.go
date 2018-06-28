package palindrome_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeNumber(t *testing.T) {
	assert.Equal(t, true, isPalindrome(121))
	assert.Equal(t, false, isPalindrome(-121))
	assert.Equal(t, false, isPalindrome(10))
}
