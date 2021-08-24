package main

import (
	"fmt"
)

func fibNumb(n int) int {
	x := 1
	y := 0
	for i := 0; i < n; i++ {
		x += y
		y = x - y
	}
	return y
}

func fibNumbsPrint(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(fibNumb(i), " ")
	}
}

func main() {
	var fibNumbs int
	for {
		fmt.Print("Enter quantity of Fibonacci numbers: ")
		fmt.Scan(&fibNumbs)
		if fibNumbs > 0 {
			fibNumbsPrint(fibNumbs)
			break
		} else {
			fmt.Println("Entered wrong value.\nPlease enter a valid integer value.\nThe value must be greater than zero")
		}
	}
}
