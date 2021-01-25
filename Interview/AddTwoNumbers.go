package Interview

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/* You are given two non-empty linked lists
representing two non-negative integers.
The digits are stored in reverse order,
and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero,
except the number 0 itself. */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0

	head := &ListNode{Val: 0}
	l3 := head

	for l1 != nil || l2 != nil {
		var l1Value int
		var l2Value int

		if l1 != nil {
			l1Value = l1.Val
		} else {

			l1Value = 0
		}

		if l2 != nil {
			l2Value = l2.Val
		} else {
			l2Value = 0
		}

		sumVar := l1Value + l2Value + carry

		carry = sumVar / 10

		lastDigit := sumVar % 10

		lX := &ListNode{Val: lastDigit}
		l3.Next = lX

		if l2 != nil {
			l2 = l2.Next
		}

		if l1 != nil {
			l1 = l1.Next
		}

		// L3 is now the previous lx
		l3 = l3.Next

	}

	if carry > 0 {
		first_node := &ListNode{Val: carry}
		l3.Next = first_node
		l3 = l3.Next
	}

	return head.Next
}


