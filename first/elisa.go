package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	x := doctor.Intro()
	fmt.Println(x)

	for {

		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		// fmt.Println(userInput) //Printing what user typed
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break

		} else {
			response := doctor.Response(userInput)
			fmt.Println(response)

		}

	}
}
