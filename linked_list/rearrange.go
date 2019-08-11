package main

import (
	"../util"
)

func findMiddleNode(head *util.LNode) *util.LNode {
	if head ==nil || head.Next == nil{
		return head
	}
	fast := head
	slow := head
	slowPre := head
	for fast != nil && fast.Next != nil{
		slowPre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	slowPre.Next = nil
	return slow
}



