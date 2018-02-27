package main

import "fmt"

func main() {
	fmt.Print("Enter number of feet: ")
	var input float64
	fmt.Scanf("%f", &input)

	meters := input * 0.3048

	fmt.Println("Number of meters:", meters)
}
