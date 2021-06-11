package lalgorithm

import "fmt"

// BinSearch 二分查找 arr 有序的
func BinSearch(arr []int, find int) int {
	min := 0
	max := len(arr) - 1
	for min <= max {
		mid := (max + min) / 2
		fmt.Println(mid)
		if arr[mid] > find {
			max = mid - 1
		} else if arr[mid] < find {
			min = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func Search() {
	arr := make([]int, 1024*1024, 1024*1024)
	for i := 0; i < 1024*1024; i++ {
		arr[i] = i + 1
	}
	id := BinSearch(arr, 1024)
	if id != -1 {
		fmt.Println(id, arr[id])
	} else {
		fmt.Println("没有找到数据")
	}
}