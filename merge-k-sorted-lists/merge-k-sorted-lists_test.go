package merge_k_sorted_lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKSortedLists(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 1

	l2 := new(ListNode)
	l2.Val = 1

	res := new(ListNode)
	res.Val = 1
	res.Next = new(ListNode)
	res.Next.Val = 1
	assert.Equal(t, res, mergeKLists([]*ListNode{l1, l2}))

	assert.Equal(t, nil, mergeKLists([]*ListNode{nil, nil}))

	assert.Equal(t, l2, mergeKLists([]*ListNode{nil, l2}))
}
