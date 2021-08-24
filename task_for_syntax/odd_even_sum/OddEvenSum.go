package main

import (
	"fmt"
)

func evenSumAndOddSum(n int) (int, int) {
	evenSum := 0
	oddSum := 0

	for i := 1; i <= 2*n; i++ {
		if (i & 1) == 0 {
			evenSum += i
		} else {
			oddSum += i
		}
	}

	return evenSum, oddSum
}

func main() {
	var numb int
	for {
		fmt.Print("Enter number: ")
		fmt.Scan(&numb)
		if numb > 0 {
			evenSum, oddSum := evenSumAndOddSum(numb)
			fmt.Println("Even sum: ", evenSum)
			fmt.Println("Odd sum: ", oddSum)
			fmt.Println("Odd and Even sum: ", oddSum+evenSum)
			break
		} else {
			fmt.Println("Entered wrong value.\nPlease enter a valid integer value.\nThe value must be greater than zero")
		}
	}
}
