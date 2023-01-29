package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
			"Saturday", "Sunday"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "Monday" {
		fmt.Println("Hard day ")
	} else if result == "Friday" {
		fmt.Println("Party HARD!!!")
	} else {
		fmt.Println("Kurwa work")
	}
}
