package page

import (
	"client/basic"
	"fmt"
	r "github.com/vasiliy-klimov-idf/openrpg_core/src/races"
	"strings"

	"github.com/manifoldco/promptui"
)

func SelectRace() string {
	basic.ClearConsole()

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\U0001F336 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "\U0001F336 {{ .Name | red | cyan }}",
		Details: `
--------- Race ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Info:" | faint }}	{{ .Description }}`,
	}

	searcher := func(input string, index int) bool {
		race := r.RaceArray[index]
		name := strings.Replace(strings.ToLower(race.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Hm... Choose your race.",
		Items:     r.RaceArray,
		Templates: templates,
		Size:      5,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return r.RaceArray[i].Name
}
