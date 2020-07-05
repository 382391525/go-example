package arr

func ArrSwap(arr []interface{}) []interface{} {
	i := 0
	j := len(arr)
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}
