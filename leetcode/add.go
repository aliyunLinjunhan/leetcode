package leetcode

 type ListNode struct {
    Val int
    Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	res := &ListNode{}
	head := res
	for ;l1 != nil || l2 != nil; {
		if l1 != nil && l2 != nil {
			res.Val += l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
			if l1 != nil || l2 != nil || res.Val >= 10{
				next := &ListNode{}
				res.Next = next
				if res.Val >= 10 {
					next.Val = 1
					res.Val = res.Val % 10
				}
				res = next
			}

		}else if l1 == nil {
			res.Val += l2.Val
			l2 = l2.Next
			if l2 != nil || res.Val >= 10{
				next := &ListNode{}
				res.Next = next
				if res.Val >= 10 {
					next.Val = 1
					res.Val = res.Val % 10
				}
				res = next
			}
		}else {
			res.Val += l1.Val
			l1 = l1.Next
			if l1 != nil || res.Val >= 10{
				next := &ListNode{}
				res.Next = next
				if res.Val >= 10 {
					next.Val = 1
					res.Val = res.Val % 10
				}
				res = next
			}
		}
	}
	return head
}

func main() {

}
