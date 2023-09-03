package main

import (
	"fmt"
	"time"
)

func cetakKelipatan(x int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Kelipatan %d: %d\n", x, x*i)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	x := 2
	go cetakKelipatan(x)

	time.Sleep(15 * time.Second)
}
