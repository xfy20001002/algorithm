package main

func mergeSort(arr []int, start, end int) {
	if start < end {
		mid := (start + end) / 2
		//17 0-8 9-16 mid=8 end=16
		//16 0-7 8-15 mid=7 end=15
		mergeSort(arr, start, mid)
		mergeSort(arr, mid+1, end)
		merge(arr, start, mid, end)
	}
}

func merge(arr []int, start, mid, end int) {
	//create two new array
	//1 2 3 s=0 mid=1 end=2
	//1 2 3 4 s=0 mid=1 end =3
	arr1 := arr[start : mid+1]
	//fmt.Println(arr1)
	arr2 := arr[mid+1 : end+1]
	//fmt.Println(arr2)
	newarr := merge2arr(arr1, arr2)
	for i, v := range newarr {
		arr[start+i] = v
	}
}

func merge2arr(arr1, arr2 []int) []int {
	newarr := make([]int, len(arr1)+len(arr2))
	i := 0
	p1, p2 := 0, 0
	//1 2 3
	//1 3 4 5
	for ; p1 < len(arr1) && p2 < len(arr2); i++ {
		if arr1[p1] < arr2[p2] {
			newarr[i] = arr1[p1]
			p1++
		} else {
			newarr[i] = arr2[p2]
			p2++
		}
	}
	if p1 == len(arr1) {
		for p2 < len(arr2) {
			newarr[i] = arr2[p2]
			p2++
			i++
		}
	} else if p2 == len(arr2) {
		for p1 < len(arr1) {
			newarr[i] = arr1[p1]
			p1++
			i++
		}
	}
	return newarr
}

/*
func main() {
	arr := []int{1, 4, 5, 6, 7, 2, 3, 9, 4, 3, 6, 2, 1, 3, 4, 10, 8}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	//test merge2arr

	//newarr := []int{1, 3, 2}
	//arr1 := []int{}
	//arr2 := []int{2}
	//merge(newarr, 0, 1, 2)
	//fmt.Println(newarr)

}
*/
