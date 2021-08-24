package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func removeSpaces(s string) string {
	return strings.Replace(s, " ", "", -1)
}

func removeSymbols(s string) string {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	s = re.ReplaceAllString(s, " ")
	return s
}

func isPalindrome(s string) bool {
	temp := s
	temp = removeSymbols(temp)
	s = removeSymbols(s)
	temp = removeSpaces(temp)
	s = removeSpaces(s)
	temp = reverse(temp)
	if strings.ToLower(s) == strings.ToLower(temp) {
		return true
	} else {
		return false
	}
}

func main() {
	var str string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter number or string:")
		if scanner.Scan() {
			str = scanner.Text()
		}
		if str != "" {
			if isPalindrome(str) {
				fmt.Println("Palindrome")
			} else {
				fmt.Println("Not Palindrome")
			}
			break
		} else {
			fmt.Println("Please enter a string or number")
		}
	}
}
