package main

import "log"

func main() {
	for n := addTwoNumbers(sToN([]int{2, 4, 3}), sToN([]int{5, 6, 4})); n.Next != nil || n.Val != 0; n = n.Next {
		log.Println(n.Val)
	}

	for n := addTwoNumbers(sToN([]int{2, 4, 3}), sToN([]int{5, 6})); n.Next != nil || n.Val != 0; n = n.Next {
		log.Println(n.Val)
	}

	for n := addTwoNumbers(sToN([]int{0}), sToN([]int{0})); n.Next != nil || n.Val != 0; n = n.Next {
		log.Println(n.Val)
	}

	for n := addTwoNumbers(sToN([]int{9, 9, 9, 9}), sToN([]int{9, 9, 9, 9, 9, 9, 9})); n.Next != nil || n.Val != 0; n = n.Next {
		log.Println(n.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	or_l := l
	carry := 0

	for l1.Next != nil || l2.Next != nil || l1.Val != 0 || l2.Val != 0 {
		v := (l1.Val + l2.Val + carry) % 10
		if l1.Val+l2.Val >= 10 {
			carry += 1
		}

		tmp := l.Val + v
		if tmp >= 10 {
			carry = 1
		}
		l.Val = tmp % 10

		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1.Val = 0
		}

		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2.Val = 0
		}

		lN := &ListNode{
			Val: carry,
		}
		l.Next = lN
		l = lN

		carry = 0
	}

	if carry != 0 {
		n := &ListNode{
			Val: 1,
		}

		l.Next = n
	}

	return or_l
}

func sToN(n []int) *ListNode {
	ln := &ListNode{
		Val: n[0],
	}
	or_ln := ln
	if len(n) == 1 {
		return ln
	}
	for i, num := range n {
		if i == 0 {
			continue
		}
		ln2 := &ListNode{
			Val: num,
		}
		ln.Next = ln2
		ln = ln2
	}

	return or_ln
}
