package main

import "fmt"

func main() {
	var a, b, t, L float32

	fmt.Print("Masukkan sisi atas: ")
	fmt.Scanln(&a)

	fmt.Print("Masukkan sisi bawah: ")
	fmt.Scanln(&b)

	fmt.Print("Masukkan tinggi: ")
	fmt.Scanln(&t)

	L = 0.5 * (a + b) * t

	fmt.Println("Luas trapesium:", L)
}
