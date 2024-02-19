package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var list []string
	list = append(list, "Eat Dinner")
	list = append(list, "Brush Teeth")
	fmt.Println("Hello, Welcome to Toodles!")

	for {
		fmt.Println("Would you like to add anything to your todo list [y/n]?")
		action := getInput()

		if action == "y" {
			fmt.Println("What would you like to add?")
			addTask := getInput()
			list = append(list, addTask)
			printList(list)
		} else if action == "n" {
			fmt.Println("Sounds good here is what I have for you today:")
			printList(list)
			break
		} else {
			fmt.Println("Incorrect Entry, Please try again")
		}

	}

}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	newTask, _ := reader.ReadString('\n')
	return strings.Replace(newTask, "\n", "", -1)
}

func printList(list []string) {

	for item := range list {
		fmt.Println(list[item])
	}
}
