package add_two_numbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := new(ListNode)

	p1 := l1
	p2 := l2
	p := res

	carry := 0

	for p1 != nil && p2 != nil {
		if carry > 0 {
			p.Val = p1.Val + p2.Val + carry
			carry = 0
		} else {
			p.Val = p1.Val + p2.Val
		}

		if p.Val >= 10 {
			carry = 1
			p.Val = p.Val - 10
		}

		p1 = p1.Next
		p2 = p2.Next
		if p1 != nil || p2 != nil || carry > 0 {
			p.Next = new(ListNode)
		}
		p = p.Next

	}

	// 如果p1长
	for p1 != nil {
		if carry > 0 {
			p.Val = p1.Val + carry
			carry = 0
		} else {
			p.Val = p1.Val
		}

		if p.Val >= 10 {
			carry = 1
			p.Val = p.Val - 10
		}

		p1 = p1.Next
		if p1 != nil || carry > 0 {
			p.Next = new(ListNode)
		}
		p = p.Next
	}

	// 如果p2长
	for p2 != nil {
		if carry > 0 {
			p.Val = p2.Val + carry
			carry = 0
		} else {
			p.Val = p2.Val
		}

		if p.Val >= 10 {
			carry = 1
			p.Val = p.Val - 10
		}

		p2 = p2.Next
		if p2 != nil || carry > 0 {
			p.Next = new(ListNode)
		}
		p = p.Next
	}

	// 如果还有carry
	if carry > 0 {
		p.Val = carry
		carry = 0
	}

	return res
}
