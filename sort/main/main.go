package main

import (
	"fmt"
)

func swap(a, b *int) {

	var c int
	c = *a
	*a = *b
	*b = c
}

// 插入前值
func InsertPre(slice []int, preIndex, rearIndex int) {

	fmt.Println(slice)
	change := slice[rearIndex]
	for i:=rearIndex-1; i>=preIndex; i-- {
		slice[i+1] = slice[i]
	}
	slice[preIndex] = change
}

// bubble sort 冒泡排序 稳定
func bubbleSort(slice []int) {

	len := len(slice)
	for i:=0; i<len; i++ {

		for j:=0; j < len-i-1; j++{
			if slice[j] > slice[j + 1] {
				swap(&slice[j], &slice[j+1])
			}
		}
	}
}

// select sort 选择排序 不稳定
func selectSort(slice []int) {

	len := len(slice)
	var minIndex int
	var min int
	for i:=0; i<len - 1; i++ {
		minIndex = i + 1
		min = slice[i + 1]
		for j:=i+2; j<len - 1; j++ {
			if slice[j] < min {
				minIndex = j
				min = slice[j]
			}
		}
		if slice[i] > min {
			swap(&slice[i], &slice[minIndex])
		}
	}
}

// insertSort 插入排序
func insertSort(slice []int) {

	len := len(slice)
	for i:=1; i<len-1; i++ {
		index := i
		for j:=i-1; j>=0; j-- {
			if(slice[j]>slice[i]){
				 index = j
			}
		}
		InsertPre(slice, index, i)
	}
}

// quickSort  快速排序
func quickSort(datas []int, start, end int) {

	if start >= end {
		// 如果start >= end 只有一个数不用排了
		return
	}
	middle := findMiddle(datas, start, end)    // 找出分割位置
	quickSort(datas, start, middle-1)	// 递归查找
	quickSort(datas, middle+1, end)
}
func findMiddle(datas []int, start, end int) int{

	temp := datas[end]
	left := start
	right := end - 1

	for true {

		for left < end && datas[left] < temp {
			left++ // 1. 从左边开始找比参照物大的值
		}
		if left == end {
			break  // 参照物是最大的
		}

		for right >= start && datas[right] > temp {
			right-- // 2. 从右边开始找比参照物小的值
		}
		// 3. 比较左右指
		if left >= right {
			// 4. 如果出现交叉, 则交换左边和参照物的值
			datas[left], datas[end] = datas[end], datas[left]
			break
		}else{
			// 5. 如果没有出现交叉，则交换左右俩边的值
			datas[left], datas[right] = datas[right], datas[left]
		}
	}
	return left
}

func void(b [3]int) {

	b[2] = 5
}

func main() {

	a := []int{20, 9, 32, 43, 2, 54, 10, 20, 30, 14, 98}
	//bubbleSort(a)
	//selectSort(a)
	//insertSort(a)
	quickSort(a, 0, len(a) - 1)

	fmt.Println(a)
}
