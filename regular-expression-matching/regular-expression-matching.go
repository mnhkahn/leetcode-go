package regular_expression_matching

import (
	"regexp/syntax"
)

func isMatch(s string, p string) bool {
	//regexp.Compile("a")
	regs, err := syntax.Parse(p, syntax.Perl)
	_, _ = regs, err

	return false
}
