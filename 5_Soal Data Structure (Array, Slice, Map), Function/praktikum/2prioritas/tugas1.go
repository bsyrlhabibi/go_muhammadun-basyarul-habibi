package main

import "fmt"

func PairSum(arr []int, target int) []int {
	// your code here
	numMap := make(map[int]int)

	for awal, number := range arr {
		complement := target - number
		if akhir, found := numMap[complement]; found {
			return []int{akhir, awal}
		}
		numMap[number] = awal
	}

	return nil
}

func main() {
	// Test cases
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}
