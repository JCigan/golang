package main

import "fmt"
import "time"
import "math/rand"

func main() {
	rand.Seed(time.Now().UnixNano())
	randomInt := random(1, 10)
	var numberName string
	switch randomInt {
	case 1:
		numberName = "One"
	case 2:
		numberName = "Two"
	case 3:
		numberName = "Three"
	case 4:
		numberName = "Four"
	case 5:
		numberName = "Five"
	case 6:
		numberName = "Six"
	case 7:
		numberName = "Seven"
	case 8:
		numberName = "Eight"
	case 9:
		numberName = "Nine"
	default:
		numberName = "Forgot to account for this number..."
	}
	fmt.Println("A random number for you is: " + numberName)
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
