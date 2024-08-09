package main


import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
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

			fmt.Println("Enter a task number from the list to remove")
			scanner.Scan()
			tasked := strings.TrimSpace(scanner.Text())

			

			//create a logic to check if the value to remove exists
			// convert the string input to integer && insert a condition to catch any error
			
			index, err := strconv.Atoi(tasked)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid integer.")
				break
			}

			//validate if the index is within the range of tasks in the slice
			if index < 1 || index >= len(tasks) {
				fmt.Println("Invalid Index, Enter a valid Number/Index")
				break
			}

			//Remove the particular element requested by slicing the slice
			tasks = append(tasks[:index], tasks[index+1:]...)
			fmt.Println("Task removed successfully....")
			fmt.Println("Updated list now is:")
			for i, v := range tasks{
				fmt.Printf("%v. is %v\n", i+1 , v)
			}
		case "lists":
			fmt.Println("Your tasks are: ")
			if len(tasks) == 0{
				fmt.Println("You have to add a task in order to get it listed ")
			} else {
				for i, v := range tasks{
					fmt.Printf("%v. is %v\n", i+1 , v)
	
				}
			}
		case "exit":
			fmt.Println("Terminating now!!")
			return
		default:
			fmt.Println("Invalid Command!!! Type any of the commands suggested")
		}
	}
	
}
