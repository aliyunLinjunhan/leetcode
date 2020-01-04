package leetcode

import (
	"fmt"
)

type RamNode struct {

	// 定义一个空闲的磁盘块类
	head int	// 表示该磁盘块的头节点
	size int	// 表示磁盘的大小为
	next *RamNode	// 下一个结点
}

type Homework struct {

	// 定义一个作业类
	no int	// 作业编号
	head int	// 占用内存块头
	tail int	// 占用内存块尾
}

type ListLink struct {

	head *RamNode
}

func(l *ListLink) Init() *ListLink {

	// 创建分区链表
	return &ListLink{nil}
}

func(l *ListLink) print() {

	fmt.Println("空闲内存分区链：")
	p := l.head
	for ;p != nil; {
		fmt.Println(*p)
		p = p.next
	}
}

func(r *RamNode) Next() *RamNode {

	if r.next != nil {
		return r.next
	}
	return nil
}

func(r *ListLink) Apply(size int, no int) *Homework {

	// 在空闲内存中申请新的空间
	var p *RamNode = r.head
	if p != nil{
		for ;p.size < size;{
			p = p.next
			if p == nil {
				fmt.Println("没能找到足够大小的空闲内存（首次适应算法）！！")
				return nil
			}
		}
	}
	h := Homework{no, p.head, p.head+size-1}
	p.head = p.head + size
	p.size = p.size - ( h.tail - h.head + 1)
	return &h
}

func(l *ListLink) Free(h *Homework) {

	// 作业释放
	head := h.head - 1// 作业头连接点
	tail := h.tail + 1 // 作业尾部连接点
	var ph *RamNode = nil
	var pt *RamNode = nil

	p := l.head
	if p == nil {
		fmt.Println("内存空闲区为空！")
		return
	}

	// 确认有没有可以进行相连接的空闲块
	for ;p != nil; {
		if p.head == tail {
			pt = p
			break
		}
		p = p.next
	}
	p = l.head
	for ;p != nil; {
		if p.head + p.size -1 == head {
			ph = p
			break
		}
		p = p.next
	}

	if pt != nil && ph != nil {

		// 前后都可以
		ph.size = ph.size + pt.size + h.tail - h.head + 1  // 重新确立前方分区大小
		ph.next = pt.next
		return
	}else if ph != nil && pt == nil {

		// 使用前面的块链接
		ph.size = ph.size + h.tail - h.head + 1
		return
	}else if pt != nil && ph == nil {

		// 使用后面的块链接
		pt.head = h.head
		pt.size = pt.size + h.tail - h.head + 1
	}else {

		newRamNode := &RamNode{h.head, h.tail - h.head + 1, nil}
		p = l.head
		var h *RamNode = nil
		for ;p.head < newRamNode.head + newRamNode.size; {
			// 找到插入点
			h = p
			p = p.next
			if p == nil {
				break
			}
		}
		if h == nil {
			l.head = newRamNode
		}else {
			h.next = newRamNode
		}
		newRamNode.next = p
		return
	}
}

func(l *ListLink) FreeBest(h *Homework) {

	// 作业最佳适应释放
	head := h.head - 1 // 作业头连接点
	tail := h.tail + 1 // 作业尾部连接点
	var ph *RamNode = nil
	var phh *RamNode = nil
	var pt *RamNode = nil
	var pth *RamNode = nil

	p := l.head
	if p == nil {
		fmt.Println("内存空闲区为空！")
		return
	}

	// 确认有没有可以进行相连接的空闲块
	for ; p != nil; {
		if p.head == tail {
			pt = p
			break
		}
		pth = p
		p = p.next
	}
	p = l.head
	for ; p != nil; {
		if p.head+p.size-1 == head {
			ph = p
			break
		}
		phh = p
		p = p.next
	}

	if pt != nil && ph != nil {

		// 前后都可以
		ph.size = ph.size + pt.size + h.tail - h.head + 1 // 重新确立前方分区大小
		ph.next = pt.next

		// 重新排序
		var p *RamNode = ph.next
		var p_h *RamNode = ph.next
		if p != nil {
			for ; p.size < ph.size; {
				p_h = p
				p = p.next
				if p == nil {
					break
				}
			}
			if phh != nil {
				phh.next = ph.next
			} else {
				l.head = ph.next
			}
			p_h.next = ph
			ph.next = p
		}
		return
	} else if ph != nil && pt == nil {

		// 使用前面的块链接
		ph.size = ph.size + h.tail - h.head + 1
		// 重新排序
		var p *RamNode = ph.next
		var p_h *RamNode = ph.next
		if p != nil {
			for ; p.size < ph.size; {
				p_h = p
				p = p.next
				if p == nil {
					break
				}
			}
			if phh != nil {
				phh.next = ph.next
			} else {
				l.head = ph.next
			}
			p_h.next = ph
			ph.next = p
		}
		return
	} else if pt != nil && ph == nil {

		// 使用后面的块链接
		pt.head = h.head
		pt.size = pt.size + h.tail - h.head + 1

		// 重新排序
		var p *RamNode = pt.next
		var p_h *RamNode = pt.next
		if p != nil {
			for ; p.size < pt.size; {
				p_h = p
				p = p.next
				if p == nil {
					break
				}
			}
			if pth != nil {
				pth.next = pt.next
			} else {
				l.head = pt.next
				p_h.next = pt
				pt.next = p
			}
			return
		} else {

			newRamNode := &RamNode{h.head, h.tail - h.head + 1, nil}
			p = l.head
			var h *RamNode = nil
			for ; p.head < newRamNode.head+newRamNode.size; {
				// 找到插入点
				h = p
				p = p.next
				if p == nil {
					break
				}
			}
			if h == nil {
				l.head = newRamNode
			} else {
				h.next = newRamNode
			}
			newRamNode.next = p
			return
		}
	}
}

func main() {

		//// 创建一个内存分区链
		//var ram *ListLink
		//ram = ram.Init()
		//// 创建一个640K的内存区
		//allRam := &RamNode{0, 640, nil}
		//ram.head = allRam
		//fmt.Println("申请一号作业：")
		//h1 := ram.Apply(130, 1)
		//fmt.Println("作业列表：")
		//fmt.Println(*h1)
		//ram.print()
		//fmt.Println()
		//
		//fmt.Println("申请二号作业：")
		//h2 := ram.Apply(60, 2)
		//fmt.Println("作业列表：")
		//fmt.Println(*h1)
		//fmt.Println(*h2)
		//ram.print()
		//fmt.Println()
		//
		//fmt.Println("申请三号作业：")
		//h3 := ram.Apply(100, 3)
		//fmt.Println("作业列表：")
		//fmt.Println(*h1)
		//fmt.Println(*h2)
		//fmt.Println(*h3)
		//ram.print()
		//fmt.Println()
		//
		//fmt.Println("释放二号作业：")
		//ram.Free(h2)
		//fmt.Println("作业列表：")
		//fmt.Println(*h1)
		//fmt.Println(*h3)
		//ram.print()
		//fmt.Println()
		//
		//fmt.Println("申请四号作业：")
		//h4 := ram.Apply(200, 4)
		//fmt.Println("作业列表：")
		//fmt.Println(*h1)
		//fmt.Println(*h3)
		//fmt.Println(*h4)
		//ram.print()
		//fmt.Println()
		//
		//fmt.Println("释放三号作业：")
		//ram.Free(h3)
		//fmt.Println("作业列表：")
		//fmt.Println(*h1)
		//fmt.Println(*h4)
		//ram.print()
		//fmt.Println()
		//
		//fmt.Println("释放一号作业：")
		//ram.Free(h1)
		//fmt.Println("作业列表：")
		//fmt.Println(*h4)
		//ram.print()
		//fmt.Println()
		//
		//fmt.Println("申请五号作业：")
		//h5 := ram.Apply(140, 5)
		//fmt.Println("作业列表：")
		//fmt.Println(*h4)
		//fmt.Println(*h5)
		//ram.print()
		//fmt.Println()
		//
		//fmt.Println("申请六号作业：")
		//h6 := ram.Apply(60, 6)
		//fmt.Println("作业列表：")
		//fmt.Println(*h4)
		//fmt.Println(*h5)
		//fmt.Println(*h6)
		//ram.print()
		//fmt.Println()
		//
		//fmt.Println("申请七号作业：")
		//h7 := ram.Apply(50, 7)
		//fmt.Println("作业列表：")
		//fmt.Println(*h4)
		//fmt.Println(*h5)
		//fmt.Println(*h6)
		//fmt.Println(*h7)
		//ram.print()
		//fmt.Println()
		//
		//fmt.Println("申请八号作业：")
		//h8 := ram.Apply(60, 8)
		//fmt.Println("作业列表：")
		//fmt.Println(*h4)
		//fmt.Println(*h5)
		//fmt.Println(*h6)
		//fmt.Println(*h7)
		//fmt.Println(*h8)
		//ram.print()
		//fmt.Println()

		// 创建一个内存分区链
		var ram *ListLink
		ram = ram.Init()
		// 创建一个640K的内存区
		allRam := &RamNode{0, 640, nil}
		ram.head = allRam
		fmt.Println("申请一号作业：")
		h1 := ram.Apply(130, 1)
		fmt.Println("作业列表：")
		fmt.Println(*h1)
		ram.print()
		fmt.Println()

		fmt.Println("申请二号作业：")
		h2 := ram.Apply(60, 2)
		fmt.Println("作业列表：")
		fmt.Println(*h1)
		fmt.Println(*h2)
		ram.print()
		fmt.Println()

		fmt.Println("申请三号作业：")
		h3 := ram.Apply(100, 3)
		fmt.Println("作业列表：")
		fmt.Println(*h1)
		fmt.Println(*h2)
		fmt.Println(*h3)
		ram.print()
		fmt.Println()

		fmt.Println("释放二号作业：")
		ram.FreeBest(h2)
		fmt.Println("作业列表：")
		fmt.Println(*h1)
		fmt.Println(*h3)
		ram.print()
		fmt.Println()

		fmt.Println("申请四号作业：")
		h4 := ram.Apply(200, 4)
		fmt.Println("作业列表：")
		fmt.Println(*h1)
		fmt.Println(*h3)
		fmt.Println(*h4)
		ram.print()
		fmt.Println()

		fmt.Println("释放三号作业：")
		ram.FreeBest(h3)
		fmt.Println("作业列表：")
		fmt.Println(*h1)
		fmt.Println(*h4)
		ram.print()
		fmt.Println()

		fmt.Println("释放一号作业：")
		ram.FreeBest(h1)
		fmt.Println("作业列表：")
		fmt.Println(*h4)
		ram.print()
		fmt.Println()

		fmt.Println("申请五号作业：")
		h5 := ram.Apply(140, 5)
		fmt.Println("作业列表：")
		fmt.Println(*h4)
		fmt.Println(*h5)
		ram.print()
		fmt.Println()

		fmt.Println("申请六号作业：")
		h6 := ram.Apply(60, 6)
		fmt.Println("作业列表：")
		fmt.Println(*h4)
		fmt.Println(*h5)
		fmt.Println(*h6)
		ram.print()
		fmt.Println()

		fmt.Println("申请七号作业：")
		h7 := ram.Apply(50, 7)
		fmt.Println("作业列表：")
		fmt.Println(*h4)
		fmt.Println(*h5)
		fmt.Println(*h6)
		fmt.Println(*h7)
		ram.print()
		fmt.Println()

		fmt.Println("申请八号作业：")
		h8 := ram.Apply(60, 8)
		fmt.Println("作业列表：")
		fmt.Println(*h4)
		fmt.Println(*h5)
		fmt.Println(*h6)
		fmt.Println(*h7)
		fmt.Println(h8)
		ram.print()
		fmt.Println()
	}