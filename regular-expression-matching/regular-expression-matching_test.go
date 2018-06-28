package regular_expression_matching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexpMatching(t *testing.T) {
	//assert.Equal(t, true, isMatch("aa", "a"))
	assert.Equal(t, true, isMatch("aa", "a*"))
}
