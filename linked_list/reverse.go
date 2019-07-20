package main

import (
	"fmt"
	"time"
)

//链表定义
type LNode struct {
	Data interface{}
	Next *LNode
}

//创建链表
func CreateNode(node *LNode, max int) {
	cur := node
	for i := 1; i < max; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}

//打印链表的方法
func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}


func main()  {
	head := &LNode{}
	fmt.Println("就地逆序")

	CreateNode(head,8)
	PrintNode("逆序前: ",head)
	s := time.Now()
	InsertReverse(head)
	e := time.Since(s)
	PrintNode("逆序后: ",head)
	fmt.Println("花费时间："+e.String())
}


//就地逆序 花费时间：281ns 时间复杂度:O(n) n空间复杂度：O(1)
func Reverse(node *LNode)  {
	if node == nil || node.Next == nil{
		return
	}
	var (
		pre *LNode
		cur *LNode
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
func RecursiveReverse(node *LNode)  {
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

func RecursiveReverseChild(node *LNode)*LNode  {
	if node == nil || node.Next == nil{
		return node
	}
	newHead := RecursiveReverseChild(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}


//插入法 花费时间：231ns 时间复杂度:O(n) n空间复杂度：O(1) 性能相对较好
func InsertReverse(node *LNode)  {
	if node == nil || node.Next == nil{
		return
	}
	var (
		cur *LNode
		next *LNode
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


