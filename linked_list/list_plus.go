package main

import (
	"../util"
	"fmt"
	"time"
)

func main() {
	fmt.Println("链表相加")
	l1, l2 := CreateNodeT()
	util.PrintNode("head1", l1)
	util.PrintNode("head2", l2)
	s := time.Now()
	addResult := Add(l1, l2)
	e := time.Since(s)
	util.PrintNode("相加后：", addResult)
	fmt.Println("花费时间：" + e.String())
}

//花费时间：754ns 时间复杂度：O(n)  空间复杂度：O(n)
func Add(h1 *util.LNode, h2 *util.LNode) *util.LNode {
	if h1 == nil || h1.Next == nil {
		return h2
	}
	if h2 == nil || h2.Next == nil {
		return h1
	}
	c := 0                      //记录进位
	sum := 0                    //记录两个节点的和
	p1 := h1.Next               //遍历h1
	p2 := h2.Next               //遍历h2
	resultHead := &util.LNode{} //相加后链表的头节点
	p := resultHead             //指向resultHead最后1个节点
	for p1 != nil && p2 != nil {
		p.Next = &util.LNode{} //指向新创建的存储相加和的节点
		sum := p1.Data.(int) + p2.Data.(int) + c
		p.Next.Data = sum % 10 //两节点相加和
		c = sum / 10           //进位
		p = p.Next
		p1 = p1.Next
		p2 = p2.Next
	}

	if p1 == nil {
		for p2 != nil {
			p.Next = &util.LNode{}
			sum = p2.Data.(int) + c
			p.Next.Data = sum % 10
			c = sum / 10
			p = p.Next
			p2 = p2.Next
		}
	}

	if p2 == nil {
		for p1 != nil {
			p.Next = &util.LNode{}
			sum = p1.Data.(int) + c
			p.Next.Data = sum % 10
			c = sum / 10
			p = p.Next
			p1 = p1.Next
		}
	}

	if c == 1 {
		p.Next = &util.LNode{}
		p.Next.Data = 1
	}

	return resultHead
}

func CreateNodeT() (l1, l2 *util.LNode) {
	l1 = &util.LNode{}
	l2 = &util.LNode{}
	cur := l1
	for i := 1; i < 7; i++ {
		cur.Next = &util.LNode{}
		cur.Next.Data = i + 2
		cur = cur.Next
	}
	cur = l2
	for i := 9; i > 4; i-- {
		cur.Next = &util.LNode{}
		cur.Next.Data = i

		cur = cur.Next
	}
	return
}
