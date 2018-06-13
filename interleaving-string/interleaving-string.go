package interleaving_string

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	if len(s1) == 0 {
		if s2 == s3 {
			return true
		} else {
			return false
		}
	}

	if len(s2) == 0 {
		if s1 == s3 {
			return true
		} else {
			return false
		}
	}

	a := s1[0]
	b := s2[0]
	c := s3[0]

	if a == b && a == c {
		if isInterleave(s1[1:], s2, s3[1:]) || isInterleave(s1, s2[1:], s3[1:]) {
			return true
		}
	} else if a == c {
		return isInterleave(s1[1:], s2, s3[1:])
	} else if b == c {
		return isInterleave(s1, s2[1:], s3[1:])
	}

	return false
}
