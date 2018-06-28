package zigzag_conversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZigZagConversion(t *testing.T) {
	assert.Equal(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
	assert.Equal(t, "PINALSIGYAHRPI", convert("PAYPALISHIRING", 4))
	assert.Equal(t, "A", convert("A", 1))
	assert.Equal(t, "ABCED", convert("ABCDE", 4))
}
