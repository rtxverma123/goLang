package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "Press Enter when done"

func main() {
	var firstNumber = 2
	var secondNumber = 5
	var subtraction = 7

	var answer int

	fmt.Println("Guess The Number Game")
	fmt.Println("----------------------")
	fmt.Println("Think of a number 1 and 10", prompt)

	reader := bufio.NewReader(os.Stdin)

	reader.ReadString('\n')

	fmt.Println("Multiply your number by ", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now Multiply your number by ", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide your number by ", prompt)
	reader.ReadString('\n')

	fmt.Println("Now Subtract ", subtraction, prompt)
	reader.ReadString('\n')

	answer = firstNumber*secondNumber - subtraction
	fmt.Println("The answer is: ", answer)
}
