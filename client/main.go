package main

import (
	"bufio"
	"client/basic"
	"client/core"
	"client/page"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/vasiliy-klimov-idf/openrpg_core/src/player"
	"os"
	"runtime"
	"time"
)

var (
	p = player.Player{}
)

func startScreen() {
	fmt.Println(basic.GameBanner)
	fmt.Println(basic.ColorBlue, basic.WelcomeText)
	fmt.Println(basic.ColorBlue, basic.Owner)
	fmt.Println(basic.ColorGreen, "Version: ", basic.AppVersion)
	fmt.Println("")
}

func startMenu() {
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
		p.Nickname = page.InputNick()

		page.SelectRaceMenu()
		p.Class.Race = core.SelectedRace

		page.SelectClassMenu()
		p.Class = core.SelectedClass

		p.CreatePlayer(p.Nickname, p.Class, p.Class.Race)
		p.PrintPlayerInfo()
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
	startMenu()

	fmt.Println("To be continues...")
	time.Sleep(30 * time.Second)
}
