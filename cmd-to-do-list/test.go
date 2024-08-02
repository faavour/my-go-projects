package main

import (
	"fmt"
	"os"
	"bufio"

)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Testing this package")

	for scanner.Scan() {
		//Read the input from the user

		input := scanner.Text()

		//process the input (print the input)
		fmt.Println("You entered:", input)

		//optionally break the loop based on some condition

		if input == "exit"{
			break
		}
	}

	// check for errors during scanning and print them
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

}