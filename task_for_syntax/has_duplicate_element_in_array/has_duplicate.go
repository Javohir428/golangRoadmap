package main

import (
	"fmt"
)

func dupCount(list []string) map[string]int {

	dupFrequency := make(map[string]int)

	for _, item := range list {
		_, exist := dupFrequency[item]
		if exist {
			dupFrequency[item] += 1
		} else {
			dupFrequency[item] = 1
		}
	}
	return dupFrequency
}

func printResult(m map[string]int) {
	for k, v := range m {
		fmt.Printf("Element : %s , Count : %d\n", k, v)
	}
}

func main() {
	var arrSize int
	for {
		fmt.Print("Enter array size: ")
		fmt.Scan(&arrSize)
		if arrSize > 0 {
			arr := make([]string, arrSize)
			fmt.Println("Enter array elements by a space: ")
			for i := 0; i < arrSize; i++ {
				fmt.Scan(&arr[i])
			}
			dupMap := dupCount(arr)
			fmt.Println("Array:", arr)
			printResult(dupMap)
			break
		} else {
			fmt.Println("Entered wrong value.\nPlease enter a valid integer value.\nThe value must be greater than zero")
		}
	}
}
