package main

import "fmt"

func main() {
	var number, num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	num = number
	if num%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

}
