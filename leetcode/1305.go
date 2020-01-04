package leetcode

import "fmt"

/**
Definition for a binary tree node. **/
 type TreeNode struct {
    Val int
     Left *TreeNode
     Right *TreeNode
}

func inOrderTravel(node *TreeNode, v *[]int) []int{

	if node == nil {
		return nil
	}
	inOrderTravel(node.Left, v)
	*v = append(*v, node.Val)
	inOrderTravel(node.Right, v)
	return *v
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {

	var a []int
	inOrderTravel(root1, &a)
	var b []int
	inOrderTravel(root2, &b)

	var r []int
	for len(a) != 0 || len(b) != 0 {
		if len(a) == 0{
			r = append(r, b...)
			break
		}
		if len(b) == 0{
			r = append(r, a...)
			break
		}
		if a[0]>b[0]{
			r = append(r, b[0])
			if len(b) >=2 {
				b = b[1:]
			}else{
				b = nil
			}
		}else{
			r = append(r, a[0])
			if len(a) >=2 {
				a = a[1:]
			}else{
				a = nil
			}
		}
	}
	return r
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
	r:= getAllElements(a,a1)
	fmt.Println(r)
}
