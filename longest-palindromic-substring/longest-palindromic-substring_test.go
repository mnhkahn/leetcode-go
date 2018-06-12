package longest_palindromic_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindromicSubString(t *testing.T) {
	assert.Equal(t, "a", longestPalindrome("a"))
	assert.Equal(t, "bab", longestPalindrome("babad"))
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
}
