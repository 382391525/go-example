package lalgorithm

import "fmt"

// QuickSort 快速排序
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

func Sort() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	//fmt.Println(QuickSort(arr))
	//fmt.Println(BubbleSort(arr))
	fmt.Println(SelectSort(arr))
}
