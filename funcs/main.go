package main

import "fmt"

func main() {
	var num1, num2, x, y int
	fmt.Print("Enter first Number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter Second Number: ")
	fmt.Scan(&num2)
	x = num1
	y = num2
	var result int
	result = max(x, y)
	fmt.Printf("max value is %d\n", result)
}

func max(a, b int) int {
	var res int
	if a > b {
		res = a
	} else {
		res = b
	}
	return res
}
