package main

import (
	"fmt"
)

func main() {
	var number, n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	n = number
	for i := 0; i < n; i++ {
		for j := 1; j < i; j++ {
			fmt.Printf("%d", j)

		}
		fmt.Println()
	}
}
