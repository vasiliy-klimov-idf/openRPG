package page

import (
	"errors"
	"github.com/manifoldco/promptui"
)

func InputNick() string {

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
