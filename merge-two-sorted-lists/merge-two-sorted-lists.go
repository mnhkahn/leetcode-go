package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func copyListNode(l *ListNode) *ListNode {
	node := new(ListNode)
	node.Val = l.Val
	node.Next = l.Next
	return node
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var res *ListNode

	a := l1
	b := l2

	if a.Val < b.Val {
		res = copyListNode(a)
		a = a.Next
	} else {
		res = copyListNode(b)
		b = b.Next
	}

	p := res

	for a != nil && b != nil {
		if a.Val < b.Val {
			p.Next = copyListNode(a)
			p = p.Next
			a = a.Next
		} else {
			p.Next = copyListNode(b)
			p = p.Next
			b = b.Next
		}
	}

	for a != nil {
		p.Next = copyListNode(a)
		p = p.Next
		a = a.Next
	}

	for b != nil {
		p.Next = copyListNode(b)
		p = p.Next
		b = b.Next
	}

	return res
}
