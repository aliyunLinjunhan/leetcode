package leetcode

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func getHeight(node *TreeNode) int{

	// 计算节点的高度
	if node == nil {
		return -1
	}
	l := getHeight(node.Left)
	r := getHeight(node.Right)
	if l > r {
		return l + 1
	}else {
		return r + 1
	}
}

func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}
	return math.Abs(float64(getHeight(root.Left) - getHeight(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func main() {

	a := &TreeNode{2, nil, nil}
	b := &TreeNode{1, nil, nil}
	c := &TreeNode{4, nil, nil}
	a.Left = b
	a.Right = c
	a1 := &TreeNode{1, nil, nil}
	b1 := &TreeNode{0, nil, nil}
	c1 := &TreeNode{3, nil, nil}
	a1.Left = b1
	a1.Right = c1
	//b.Left = a1
	r := getHeight(a)
	ib := isBalanced(a)
	fmt.Println(r, ib)
}
