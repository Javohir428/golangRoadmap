package main

import (
	"fmt"
)

func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if ((i % 5) == 0) && ((i % 7) == 0) {
			fmt.Println("fizzbuzz")
		} else if (i % 5) == 0 {
			fmt.Println("fizz")
		} else if (i % 7) == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	var numb int
	for {
		fmt.Print("Enter number: ")
		fmt.Scan(&numb)
		if numb > 0 {
			fizzBuzz(numb)
			break
		} else {
			fmt.Println("Entered wrong value.\nPlease enter a valid integer value.\nThe value must be greater than zero")
		}
	}
}
