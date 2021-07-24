package main

import "fmt"

func main() {
	var n = []int{12, 56, 76, 98, 19, 13}                                //Dynamic arrays
	var x = [5]string{"Ankita", "Prerna", "Smarika", "Nitesh", "Sriram"} //Static array

	for i := 0; i < 6; i++ {
		fmt.Println("The array value is: ", n[i])

	}

	var j int
	for j = 0; j <= 5; j++ {
		fmt.Println("The value of string array is: ", x[j])
	}

	var y [10]int
	var k int
	for k = 0; k < 10; k++ {
		y[k] = k + 100
	}
	var p int
	for p = 0; p < 10; p++ {
		fmt.Printf("Element[%d] = %d\n", p, y[p])
	}

}
