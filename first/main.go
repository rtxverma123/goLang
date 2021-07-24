package main

import "fmt"

func maain() { //Main function dont have any paramteres
	fmt.Println("Hello World")
	fmt.Println("HI! I am Tamish") //It will not make it go to new line
	sayHelloWorld("Hi from functional programming")

}

func sayHelloWorld(somesaid string) {
	fmt.Println(somesaid)
}
