package main

import (
	"../util"
	"fmt"
	"time"
)


func main()  {
	head := &util.LNode{}
	fmt.Println("就地逆序")

	util.CreateNode(head,8)
	util.PrintNode("逆序前: ",head)
	s := time.Now()
	InsertReverse(head)
	e := time.Since(s)
	util.PrintNode("逆序后: ",head)
	fmt.Println("花费时间："+e.String())
}


//就地逆序 花费时间：281ns 时间复杂度:O(n) n空间复杂度：O(1)
func Reverse(node *util.LNode)  {
	if node == nil || node.Next == nil{
		return
	}
	var (
		pre *util.LNode
		cur *util.LNode
	)
	next := node.Next
	for next != nil{
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	node.Next = pre
}



//递归法 花费时间：257ns 时间复杂度:O(n) n空间复杂度：O(1)
func RecursiveReverse(node *util.LNode)  {
	firstNode := node.Next
	/*
		内部调用相对耗时
		var reverse func(node *LNode)*LNode
		reverse = func(node *LNode)*LNode {
			if node == nil || node.Next == nil{
				return node
			}
			newHead := reverse(node.Next)
			node.Next.Next = node
			node.Next = nil
			return newHead
		}*/

	newHead := RecursiveReverseChild(firstNode)
	node.Next = newHead
}

func RecursiveReverseChild(node *util.LNode)*util.LNode  {
	if node == nil || node.Next == nil{
		return node
	}
	newHead := RecursiveReverseChild(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}


//插入法 花费时间：231ns 时间复杂度:O(n) n空间复杂度：O(1) 性能相对较好
func InsertReverse(node *util.LNode)  {
	if node == nil || node.Next == nil{
		return
	}
	var (
		cur *util.LNode
		next *util.LNode
	)
	cur = node.Next.Next
	node.Next.Next = nil
	for cur != nil{
		next = cur.Next
		cur.Next = node.Next
		node.Next = cur
		cur = next
	}
}


