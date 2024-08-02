package main


import (
	"fmt"
	"os"
	"bufio"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	// texts := []string{}

	for scanner.Scan() {
		fmt.Println("Enter command (add, remove, complete, lists, or exit):")
		//Read the input from the user

		input := scanner.Text()

		//process the input (print the input)
		fmt.Println("You entered:", input)

		//optionally break the loop based on some condition

		if input == "exit"{
			break
		}
	}
}
