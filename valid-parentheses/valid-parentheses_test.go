package valid_parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidParentheses(t *testing.T) {
	assert.Equal(t, true, isValid(""))
	assert.Equal(t, true, isValid("{}[]{}"))
	assert.Equal(t, false, isValid("(]"))
	assert.Equal(t, false, isValid("([)]"))
	assert.Equal(t, true, isValid("{[]}"))
}
