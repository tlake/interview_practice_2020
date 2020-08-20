package addtwonumbers

/*
LeetCode Difficulty: Medium

You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

*/

/*
Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}
/*

/*
Reflection:

Initial implementation took 15 minutes.
Sorting out edge cases took another 15, though.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	node1, node2 := l1, l2
	carry := 0

	for node1 != nil || node2 != nil {
		if node1 == nil {
			node1 = &ListNode{Val: 0}
		}

		if node2 == nil {
			node2 = &ListNode{Val: 0}
		}

		result := node1.Val + node2.Val + carry
		carry = 0

		if result > 9 {
			carry = 1
			result = result - 10
		}

		newNode := &ListNode{Val: result}

		if tail != nil {
			tail.Next = newNode
		}

		tail = newNode

		if head == nil {
			head = newNode
		}

		node1, node2 = node1.Next, node2.Next
	}

	if carry > 0 {
		newNode := &ListNode{Val: carry}
		tail.Next = newNode
		tail = newNode
	}

	return head
}
