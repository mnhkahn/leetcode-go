package add_two_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 2
	l1.Next = new(ListNode)
	l1.Next.Val = 4
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 3

	l2 := new(ListNode)
	l2.Val = 5
	l2.Next = new(ListNode)
	l2.Next.Val = 6
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 4

	res := new(ListNode)
	res.Val = 7
	res.Next = new(ListNode)
	res.Next.Val = 0
	res.Next.Next = new(ListNode)
	res.Next.Next.Val = 8

	assert.Equal(t, res, addTwoNumbers(l1, l2))
}

func TestFiveFive(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 5

	l2 := new(ListNode)
	l2.Val = 5

	res := new(ListNode)
	res.Val = 0
	res.Next = new(ListNode)
	res.Next.Val = 1

	assert.Equal(t, res, addTwoNumbers(l1, l2))
}

func TestAddZero(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = new(ListNode)
	l1.Next.Val = 8

	l2 := new(ListNode)
	l2.Val = 0

	res := new(ListNode)
	res.Val = 1
	res.Next = new(ListNode)
	res.Next.Val = 8

	assert.Equal(t, res, addTwoNumbers(l1, l2))
}

func TestAddTo100(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 1

	l2 := new(ListNode)
	l2.Val = 9
	l2.Next = new(ListNode)
	l2.Next.Val = 9

	res := new(ListNode)
	res.Val = 0
	res.Next = new(ListNode)
	res.Next.Val = 0
	res.Next.Next = new(ListNode)
	res.Next.Next.Val = 1

	assert.Equal(t, res, addTwoNumbers(l1, l2))
}
