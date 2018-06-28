package merge_two_sorted_lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoSortedLists(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 1

	l2 := new(ListNode)
	l2.Val = 1

	res := new(ListNode)
	res.Val = 1
	res.Next = new(ListNode)
	res.Next.Val = 1
	assert.Equal(t, res, mergeTwoLists(l1, l2))

	assert.Equal(t, nil, mergeTwoLists(nil, nil))
}
