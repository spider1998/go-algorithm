package main

import (
	"fmt"
	"strconv"
)

type rbnode struct {
	key, value  interface{}
	red         bool
	left, right *rbnode
}

type RBTree struct {
	root   *rbnode
	length int
	less   func(a, b interface{}) bool
}

func NewRBTree(less func(a, b interface{}) bool) *RBTree {
	return &RBTree{
		less: less,
	}
}

func isRed(node *rbnode) bool {
	return node != nil && node.red
}

func newRBNode(k, v interface{}) *rbnode {
	return &rbnode{
		key:   k,
		value: v,
	}
}

func main() {
	tree := NewRBTree(func(a, b interface{}) bool {
		if a.(int) < b.(int) {
			return true
		} else {
			return false
		}
	})
	var i int
	var ok bool
	node := newRBNode(0, 000)
	i++
	for i < 100 {
		node, ok = tree.insert(node, i, "元素"+strconv.Itoa(i))
		if !ok {
			fmt.Println("---------------------")
		}
		i++
	}
	tree.root = node
	n := tree.search(76)
	fmt.Println(n)

}

func rotateLeft(node *rbnode) *rbnode {
	rightChild := node.right
	node.right = rightChild.left
	rightChild.left = node
	rightChild.red = node.red
	node.red = true
	return rightChild
}

func rotateRight(node *rbnode) *rbnode {
	leftChild := node.left
	node.left = leftChild.right
	leftChild.right = node
	leftChild.red = node.red
	node.red = true
	return leftChild
}

func colorFlip(node *rbnode) *rbnode {
	node.red = !node.red
	if node.left != nil {
		node.left.red = !node.left.red
	}
	if node.right != nil {
		node.right.red = !node.right.red
	}
	return node
}

func fixUp(node *rbnode) *rbnode {
	if isRed(node.right) { //遇到向右倾斜的红边，将其向左旋转
		node = rotateLeft(node)
	}
	if isRed(node.left) && isRed(node.left.left) { //遇到两条红边在一条路径上，对其进行平衡
		node = rotateRight(node)
	}
	if isRed(node.left) && isRed(node.right) { //分解4-node
		node = colorFlip(node)
	}
	return node
}

func (r *RBTree) search(k interface{}) *rbnode {
	node := r.root
	for {
		if node == nil {
			return nil
		}
		if r.less(k, node.key) {
			node = node.left
		} else if r.less(node.key, k) {
			node = node.right
		} else {
			return node
		}
	}
}

func (r *RBTree) insert(node *rbnode, k, v interface{}) (*rbnode, bool) {
	ok := false
	if node == nil {
		return &rbnode{key: k, value: v, red: true}, true
	}
	if r.less(k, node.key) {
		node.left, ok = r.insert(node.left, k, v)
	} else if r.less(node.key, k) {
		node.right, ok = r.insert(node.right, k, v)
	} else {
		node.value = v
		ok = true
	}
	return fixUp(node), ok
}

func moveRedRight(node *rbnode) *rbnode {
	node = colorFlip(node)
	if node.left != nil && isRed(node.left.left) {
		node = rotateRight(node)
		node = colorFlip(node)
	}
	return node
}

func moveRedLeft(node *rbnode) *rbnode {
	node = colorFlip(node)
	if isRed(node.right.left) {
		node.right = rotateRight(node.right)
		node = rotateLeft(node)
		node = colorFlip(node)
	}
	return node
}

func deleteMax(node *rbnode) *rbnode {
	if isRed(node.left) {
		node = rotateRight(node)
	}
	if node.right == nil {
		return nil
	}
	if !isRed(node.right) && !isRed(node.right.left) {
		node = moveRedRight(node)
	}
	node.right = deleteMax(node.right)
	return fixUp(node)
}

func deleteMin(node *rbnode) *rbnode {
	if node.left == nil {
		return nil
	}
	if !isRed(node.left) && !isRed(node.left.left) {
		node = moveRedLeft(node)
	}
	node.left = deleteMin(node.left)
	return fixUp(node)
}

func (r *RBTree) delete(node *rbnode, k interface{}) (*rbnode, bool) {
	ok := false
	if r.less(k, node.key) {
		if r.root.left != nil {
			if !isRed(node.left) && !isRed(node.left.left) {
				node = moveRedLeft(node)
			}
			node.left, ok = r.delete(node.left, k)
		}
	} else {
		if isRed(node.left) {
			node = rotateRight(node)
		}
		if !r.less(k, node.key) && !r.less(node.key, k) && node.right == nil {
			return nil, true
		}
		if node.right != nil {
			if !isRed(node.right) && !isRed(node.right.left) {
				node = moveRedRight(node)
			}
			if !r.less(k, node.key) && !r.less(node.key, k) {
				smallest := min(node.right)
				node.key = smallest.key
				node.value = smallest.value
				node.right = deleteMin(node.right)
				ok = true
			} else {
				node.right, ok = r.delete(node.right, k)
			}
		}
	}
	return fixUp(node), ok
}

func min(node *rbnode) *rbnode {
	for node.left != nil {
		node = node.left
	}
	return node
}
