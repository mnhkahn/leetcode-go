package longest_valid_parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestValidParentheses(t *testing.T) {
	//assert.Equal(t, 2, longestValidParentheses("(()"))
	//assert.Equal(t, 4, longestValidParentheses(")()())"))
	//assert.Equal(t, 2, longestValidParentheses("()(()"))
	//assert.Equal(t, 4, longestValidParentheses(")()())()()("))
	assert.Equal(t, 4, longestValidParentheses("(()))())("))
}
