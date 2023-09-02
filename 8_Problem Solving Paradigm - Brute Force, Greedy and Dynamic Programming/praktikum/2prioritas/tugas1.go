package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	// your code here
	costs := make([]int, len(jumps))
	for i := range costs {
		costs[i] = math.MaxInt32
	}

	costs[0] = 0

	for i := 0; i < len(jumps)-1; i++ {
		for j := 1; j <= 2; j++ {
			if i+j < len(jumps) {
				cost := costs[i] + int(math.Abs(float64(jumps[i])-float64(jumps[i+j])))
				if cost < costs[i+j] {
					costs[i+j] = cost
				}
			}
		}
	}

	return costs[len(jumps)-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
