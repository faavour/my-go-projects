package main


import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tasks := []string{}

	for {
		fmt.Println("Enter command (add, remove, lists, or exit):")
		//Read the input from the user

		scanner.Scan()

		output := strings.TrimSpace(scanner.Text())
		switch output {
		case "add":
			
			fmt.Println("Enter task you wish to add")
			scanner.Scan()
			task := strings.TrimSpace(scanner.Text())
			tasks = append(tasks, task)
			fmt.Println(" ==> Task added successfully <==")
		
		case "remove":

			fmt.Println("To remove a task you'd love to add, type `lists` to list it")
			remove := len(tasks)
			fmt.Println(remove)
			

		case "lists":
			fmt.Println("Your tasks are: ")
			for i, v := range tasks{
				fmt.Printf("task %v is %v\n", i , v)

			}
		case "exit":
			fmt.Println("Terminating now!!")
			return
		default:
			fmt.Println("Invalid Command!!! Type any of the commands suggested")
		}
	}
	
}
