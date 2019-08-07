//从无序链表中移除重复项
package main

import (
	"../util"
	"fmt"
	"time"
)

func main() {

	head := &util.LNode{
		Next: &util.LNode{
			Data: 1,
			Next: &util.LNode{
				Data: 3,
				Next: &util.LNode{
					Data: 1,
					Next: &util.LNode{
						Data: 5,
						Next: &util.LNode{
							Data: 5,
							Next: &util.LNode{
								Data: 7,
							},
						},
					},
				},
			},
		},
	}
	fmt.Println("删除重复节点")
	fmt.Println("-------删除-------")
	util.PrintNode("删除前：", head)
	s := time.Now()
	RemoveDupRecursion(head)
	e := time.Since(s)
	util.PrintNode("删除后：", head)
	fmt.Println("花费时间：" + e.String())

}

//顺序删除 花费时间：774ns 时间复杂度O(n^2) 空间复杂度O(1)
func RemoveDup(head *util.LNode) {
	if head == nil || head.Next == nil {
		return
	}
	outerCur := head.Next
	var (
		innerCur *util.LNode
		innerPre *util.LNode
	)
	for ; outerCur != nil; outerCur = outerCur.Next {
		for innerCur, innerPre = outerCur.Next, outerCur; innerCur != nil; {
			if outerCur.Data == innerCur.Data {
				innerPre.Next = innerCur.Next
				innerCur = innerCur.Next
			} else {
				innerPre = innerCur
				innerCur = innerCur.Next
			}
		}
	}
}

func removeDupRecursionChild(head *util.LNode) *util.LNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pointer *util.LNode
	cur := head
	head.Next = removeDupRecursionChild(head.Next)
	pointer = head.Next
	for pointer != nil {
		if head.Data == pointer.Data {
			cur.Next = pointer.Next
			pointer = cur.Next
		} else {
			pointer = pointer.Next
			cur = cur.Next
		}
	}
	return head
}

//递归删除 花费时间：861ns 时间复杂度O(n^2) 空间复杂度O(1) 效率比顺序删除低
func RemoveDupRecursion(head *util.LNode) {
	if head == nil {
		return
	}
	head.Next = removeDupRecursionChild(head.Next)
}
