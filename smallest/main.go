package main

import "fmt"

func main() {
	var min int
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	for index, value := range x {
		if index == 0 || value < min {
			min = value
		}
	}
	fmt.Println("Smallest number:", min)
}
