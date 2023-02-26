package main

import (
	"client/basic"
	"client/core"
	"client/page"

	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/vasiliy-klimov-idf/openrpg_core/src/story"
	"os"
	"runtime"
	"time"
)

func startScreen() {
	fmt.Println(basic.GameBanner)
	fmt.Println(basic.ColorBlue, basic.WelcomeText)
	fmt.Println(basic.ColorBlue, basic.Owner)
	fmt.Println(basic.ColorGreen, "Version: ", basic.AppVersion)
	fmt.Println("")
}

func StartMenu() {
	var MenuList []string
	SaveFileExist := false
	if SaveFileExist == true {
		MenuList = []string{"New Game", "Continua", "Load Save Game", "Help", "Exit"}
	} else {
		MenuList = []string{"New Game", "Load Save Game", "Help", "Exit"}
	}
	prompt := promptui.Select{
		Label: "Menu",
		Items: MenuList,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "New Game" {
		story.PlayerName = page.InputNick()

		page.SelectRaceMenu()
		page.SelectClassMenu()
		core.InitPlayer()
		fmt.Println("User created")
		time.Sleep(5 * time.Second)
		basic.ClearConsole()
	} else if result == "Load Save Game" {
		fmt.Println("To be continue ...")
	} else if result == "Help" {
		fmt.Println("Good Luck! :)")
	} else {
		fmt.Println("See you later)")
		os.Exit(0)
	}
}

func main() {

	basic.TerminalCheck()
	fmt.Println("OS", runtime.GOOS)
	startScreen()
	StartMenu()

	page.TestTextPage()

	fmt.Println("To be continues...")
	time.Sleep(30 * time.Second)
}
