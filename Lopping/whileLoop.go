package main

import (
	"fmt"
	"math/rand"
	"time"
)

func maisn() {
	rand.Seed(time.Now().UnixNano())

	i := 1000

	for i > 100 {
		i = rand.Intn(1000) + 1
		fmt.Println("i is ", i)
		if i > 100 {
			fmt.Println("i is ", i, "so loop keeps going")
		}
	}

	fmt.Println("Got", i, "broke out of the loop")

}
