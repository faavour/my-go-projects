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

			fmt.Println("Enter a task name from the tasks added which you want to remove")
			scanner.Scan()
			tasked := strings.TrimSpace(scanner.Text())

			

			//create a logic to check if the value to remove exists
			index := -1

			for i, v := range tasks{
				if v == tasked{
					index = i
					break
				}

			}

			//Next we can check if the value requested existed
			if index != -1{
				//Remove the particular element requested by slicing the slice

				tasks = append(tasks[:index], tasks[index+1:]...)
				fmt.Println("Updated list now is:")
				for i, v := range tasks{
					fmt.Printf("task %v is %v\n", i+1 , v)
				}
			}
		case "lists":
			fmt.Println("Your tasks are: ")
			if len(tasks) == 0{
				fmt.Println("You have to add a task in order to get it listed ")
			} else {
				for i, v := range tasks{
					fmt.Printf("task %v is %v\n", i+1 , v)
	
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
