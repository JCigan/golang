package main

import "fmt"

func main() {
	fmt.Print("Enter a temperature Fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	celsius := (input - 32) * (5 / 9)

	fmt.Println("Degrees Celsius:", celsius)
}
