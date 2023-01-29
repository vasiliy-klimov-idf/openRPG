package main

import (
	"OpenRPG/basic"
	"bufio"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"runtime"
)

func startScreen() {
	fmt.Println(basic.GameBanner)
	fmt.Println(basic.ColorBlue, basic.WelcomeText)
	fmt.Println(basic.ColorBlue, basic.Owner)
	fmt.Println(basic.ColorGreen, "Version: ", basic.AppVersion)

	fmt.Println("")

}

func menu() {
	prompt := promptui.Select{
		Label: "Menu",
		Items: []string{"New Game", "Load Save Game", "Help", "Exit"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "New Game" {
		fmt.Println("New Game")
	} else if result == "Load Save Game" {
		fmt.Println("To be continue ...")
	} else if result == "Help" {
		fmt.Println("Good Luck! :)")
	} else {
		fmt.Println("See you later)")
		os.Exit(0)
	}
}

func loginPage() {

	fmt.Print("Login: ")
	reader := bufio.NewReader(os.Stdin)
	login, _ := reader.ReadString('\n')
	login = login //заглушка

	fmt.Printf("Password: ")
	pass, _ := basic.GetPasswd()
	// Do something with pass
	fmt.Println(string(pass))
}
func main() {
	basic.TerminalCheck()
	fmt.Println("OS", runtime.GOOS)
	startScreen()
	loginPage()
	menu()
}
