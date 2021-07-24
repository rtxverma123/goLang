package main

import "fmt"

func main() {
	var myMap map[string]int
	myMap = map[string]int{}
	myMap["a"] = 100
	myMap["b"] = 200
	myMap["c"] = 300
	myMap["d"] = 500
	myMap["e"] = 600
	fmt.Println(myMap)

	//Delete value from map
	delete(myMap, "d")

	fmt.Println("After deletion: ", myMap)

}
