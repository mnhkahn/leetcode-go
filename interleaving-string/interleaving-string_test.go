package interleaving_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInterleave(t *testing.T) {
	assert.Equal(t, true, isInterleave("c", "ca", "cac"))
	assert.Equal(t, true, isInterleave("a", "b", "ab"))
	assert.Equal(t, true, isInterleave("a", "b", "ab"))
	assert.Equal(t, true, isInterleave("bcc", "dbbca", "dbbcbcac"))
	assert.Equal(t, true, isInterleave("aabcc", "dbbca", "aadbbcbcac"))
}
