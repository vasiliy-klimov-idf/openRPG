package basic

import (
	"fmt"
	"golang.org/x/term"
	"os"
	"time"
)

func TerminalCheck() {
	if term.IsTerminal(0) {
		println("in a term")
	} else {
		println("not in a term")
	}
	width, height, err := term.GetSize(1)
	if err != nil {
		return
	}
	if width <= 150 && height <= 80 {

		fmt.Println("Terminal have a small size. Open in Terminal in FullScreen Mode.")
		os.Exit(1)
	}
	println("width:", width, "height:", height)
}

func ClearConsole() {
	fmt.Print("\033[H\033[2J")
}

func typingEffect(str string) {
	chars := []rune(str)
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		time.Sleep(100 * time.Millisecond)
		print(char)
	}
	fmt.Println()
}
