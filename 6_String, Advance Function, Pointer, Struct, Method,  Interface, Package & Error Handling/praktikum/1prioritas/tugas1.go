package main

import "fmt"

type Car struct {
	CarType string
	Fuelin  float64
}

func (c Car) Distance() (string, float64) {
	fuelFilled := 1.5
	estDistance := c.Fuelin * fuelFilled
	return c.CarType, estDistance
}

func main() {
	c := Car{
		CarType: "SUV",
		Fuelin:  10.5,
	}

	carType, estDistance := c.Distance()

	fmt.Printf("car type: %s , est. distance: %.2f\n", carType, estDistance)
}
