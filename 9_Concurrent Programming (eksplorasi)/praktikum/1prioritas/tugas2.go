package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for i := 1; i <= 10; i++ {
			channel <- i * 3
		}
	}()

	for kelipatanTiga := range channel {
		fmt.Println(kelipatanTiga)
	}
}
