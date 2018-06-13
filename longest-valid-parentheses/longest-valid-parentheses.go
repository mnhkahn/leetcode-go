package longest_valid_parentheses

func longestValidParentheses(s string) int {
	sk := new(stack)

	for i, b := range s {
		if b == ')' && sk.top() == '(' {
			sk.pop()
		} else {
			sk.push(byte(b), i)
		}
	}

	if len(sk.pos) == 0 {
		return len(s)
	}

	pre := sk.pos[0]
	res := sk.pos[0]
	for i := 1; i < len(sk.pos); i++ {
		pos := sk.pos[i]
		if pos > pre {
			tmp := pos - pre - 1
			if tmp > res {
				res = tmp
			}
		}
		pre = pos
	}

	tmp := len(s) - sk.pos[len(sk.pos)-1] - 1
	if tmp > res {
		res = tmp
	}

	return res
}

type stack struct {
	b   []byte
	pos []int
}

func (s *stack) push(b byte, pos int) {
	s.b = append(s.b, b)
	s.pos = append(s.pos, pos)
}

func (s *stack) pop() byte {
	if len(s.b) > 0 {
		res := s.top()
		s.b = s.b[0 : len(s.b)-1]
		s.pos = s.pos[0 : len(s.pos)-1]

		return res
	}

	return 0
}

func (s *stack) top() byte {
	if len(s.b) > 0 {
		return s.b[len(s.b)-1]
	}
	return 0
}
