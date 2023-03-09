package leetcode

import "fmt"

/**
 * Definition for singly-linked list
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(a []int) *ListNode {
	var head, p *ListNode = nil, nil
	for _, v := range a {
		temp := new(ListNode)
		temp.Val = v
		if head == nil {
			head = temp
			p = temp
		} else {
			p.Next = temp
			p = p.Next
		}
	}
	return head
}

func (l *ListNode) Len() int {
	lenth := 0
	for p := l; p != nil; p = p.Next {
		lenth++
	}
	return lenth
}

func (l *ListNode) Print() {
	slice := []int{}
	for p := l; p != nil; p = p.Next {
		slice = append(slice, p.Val)
	}
	fmt.Println(slice)
}
