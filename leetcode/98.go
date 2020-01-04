package leetcode

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func transSlice(node *TreeNode, v *[]int) {

	if node == nil {
		return
	}
	transSlice(node.Left, v)
	*v = append(*v, node.Val)
	transSlice(node.Right, v)
}

func isValidBST(root *TreeNode) bool {

	var slice []int
	transSlice(root, &slice)
	len := len(slice)
	for i:=0; i<len-1; i++ {
		if slice[i] >= slice[i+1] {
			return false
		}
	}
	return true
}

func main() {

	a := &TreeNode{1, nil, nil}
	b := &TreeNode{1, nil, nil}
	//c := &TreeNode{4, nil, nil}
	a.Left = b
	//a.Right = c
	a1 := &TreeNode{1, nil, nil}
	b1 := &TreeNode{0, nil, nil}
	c1 := &TreeNode{3, nil, nil}
	a1.Left = b1
	a1.Right = c1
	//b.Left = a1

	var r []int
	transSlice(a, &r)
	rb := isValidBST(a)
	fmt.Println(r)
	fmt.Println(rb)
}
