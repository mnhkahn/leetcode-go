package merge_k_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func copyListNode(l *ListNode) *ListNode {
	if l == nil {
		return l
	}
	node := new(ListNode)
	node.Val = l.Val
	node.Next = l.Next
	return node
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var res *ListNode

	ps := make([]*ListNode, len(lists))
	for i, l := range lists {
		ps[i] = copyListNode(l)
	}

	p := res

	for true {
		nilCnt := 0

		lowest := 0
		index := 0
		for i, pps := range ps {
			if pps == nil {
				nilCnt++
			} else {
				lowest = i
				break
			}
			index = i
		}

		for i := index + 1; i < len(ps); i++ {
			pps := ps[i]

			if pps == nil {
				nilCnt++
				continue
			}

			if ps[lowest].Val > pps.Val {
				lowest = i
			}
		}

		if nilCnt == len(lists) {
			break
		}

		if res == nil {
			p = ps[lowest]
			res = p
		} else {
			p.Next = ps[lowest]
			p = p.Next
		}

		ps[lowest] = ps[lowest].Next
	}

	return res
}
