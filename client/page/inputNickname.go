package page

import (
	"client/basic"
	"errors"
	"github.com/manifoldco/promptui"
)

func InputNick() string {
	basic.ClearConsole()
	validate := func(input string) error {
		if len(input) < 3 {
			return errors.New("username must have more than 3 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Username",
		Validate: validate,
		Default:  "",
	}

	result, _ := prompt.Run()
	return result
}
