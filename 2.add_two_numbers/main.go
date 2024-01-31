package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	plus := 0 // 进位
	result := new(ListNode)
	head := result

	for {
		if l1 == nil && l2 == nil && plus == 0 {
			break
		}
		l1Val := 0
		l2Val := 0
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		val := l1Val + l2Val + plus
		plus = val / 10
		val = val % 10
		result.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		result = result.Next
	}
	return head.Next
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	r := addTwoNumbers(l1, l2)
	fmt.Println(l1.Val)
	fmt.Println(l2.Val)
	fmt.Println(r)
}
