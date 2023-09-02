package main

import (
	"fmt"
	"strconv"
)

func binaryRepresentation(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		binStr := strconv.FormatInt(int64(i), 2)
		binInt, _ := strconv.Atoi(binStr)
		ans[i] = binInt
	}
	return ans
}

func main() {
	var n int
	n = 2
	output := binaryRepresentation(n)
	fmt.Println(output) // [0 1 10]

	n = 3
	output = binaryRepresentation(n)
	fmt.Println(output) // [0 1 10 11]
}
