package main

import "fmt"

// Slice is just like array but with slice we can use inbuilt functions

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slc := arr[1:4]       //Slice like list
	slc = append(slc, 45) //Add a new value to array
	fmt.Println(slc, len(slc), arr)

}
