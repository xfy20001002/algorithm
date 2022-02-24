package main

//求两数之和
func SumOfTwo(arr []int, target int) []int {
	//hash table
	var hashTable = make(map[int]int)
	for index, value := range arr {
		if val, ok := hashTable[target-value]; ok {
			return []int{index, val}
		} else {
			hashTable[value] = index
		}
	}
	return []int{}
}
