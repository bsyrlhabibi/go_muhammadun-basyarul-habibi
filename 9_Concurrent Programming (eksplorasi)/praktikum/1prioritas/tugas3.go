package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 10)

	go func() {
		defer close(channel)
		for i := 3; i < 100; i += 3 {
			channel <- i
		}
	}()

	time.Sleep(5 * time.Second)

	for i := range channel {
		fmt.Println(i)
	}
}
