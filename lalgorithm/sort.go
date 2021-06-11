package lalgorithm

import (
	"fmt"
	"sort"
)

// QuickSort 快速排序
//选一个基准值 把数据分为左右两边,左边的小于基准值右边的大于基准值,然后将两边的数据递归的重复,最后将排好序的数据合并  nlgn
func QuickSort(req []int) []int {
	if len(req) <= 1 {
		return req
	}
	//选第一个值为基准值
	num := req[0]
	//比基准值小
	min := make([]int, 0, 0)
	//比基准值大
	max := make([]int, 0, 0)
	//基准值一样大
	mid := make([]int, 0, 0)
	mid = append(mid, num)
	for i := 1; i < len(req); i++ {
		if req[i] < num {
			min = append(min, req[i])
		} else if req[i] > num {
			max = append(max, req[i])
		} else {
			mid = append(mid, req[i])
		}
	}
	min = QuickSort(min)
	max = QuickSort(max)
	return append(append(min, mid...), max...)
}

// BubbleSort 冒泡排序
// 数组的前一个数和后一个数比较大小,将最大(小)的那个值冒泡到最后
//两个for循环处理 n2
func BubbleSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

// SelectSort 选择排序
//从未排序的数据中,选择最大或最小值很当前值交换.
//有点类似冒泡
func SelectSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

// InsertSort 插入排序
//前部是排好序的,每次取未排好序的第一个往前面排好的数据中比值然后插入
func InsertSort(data []int) []int {
	// i是添加新的排序元素
	for i := 0; i < len(data); i++ {
		tmp := data[i]
		for j := i; j >= 0; j-- {
			// 放到第一个
			if j == 0 {
				data[0] = tmp
				break
			}
			if data[j-1] > tmp {
				data[j] = data[j-1]
			} else {
				// 放到最后一个
				data[j] = tmp
				break
			}
		}
	}
	return data
}

//归并排序
//堆排序

func Sort() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	//fmt.Println(QuickSort(arr))
	//fmt.Println(BubbleSort(arr))
	fmt.Println(InsertSort(arr))
}

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}