package longest_common_prefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, "", longestCommonPrefix([]string{"dog", "racecar", "car"}))
	assert.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	assert.Equal(t, "", longestCommonPrefix(nil))
	assert.Equal(t, "", longestCommonPrefix([]string{}))
}
