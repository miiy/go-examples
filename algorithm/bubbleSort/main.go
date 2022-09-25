package bubbleSort

func bubbleSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}
