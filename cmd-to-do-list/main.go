package main


import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tasks := []string{}

	for {
		fmt.Println("Enter command (add, remove, complete, lists, or exit):")
		//Read the input from the user

		input := scanner.Scan()

		//process the input (print the input)
		fmt.Println("You entered:", input)

		//optionally break the loop based on some condition

		if !input {
			break
		}

		output := scanner.Text()
		switch output {
		case "add":
			
			fmt.Println("Task added Successfully")
			task := scanner.Text()
			tasks = append(tasks, task)
			fmt.Println(tasks)
		}
	}
	
}
