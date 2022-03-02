package main

func binary_find(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	mid := (start + end) / 2
	for start <= end {
		if arr[mid] == target {
			return mid
		} else {
			if arr[mid] > target {
				end = mid - 1
			} else {
				start = mid + 1
			}
			mid = (start + end) / 2
		}
	}
	return -1
}

/*
func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13}
	res := binary_find(arr, 7)
	res = binary_find(arr, 8)
	fmt.Println(res)
}
*/
