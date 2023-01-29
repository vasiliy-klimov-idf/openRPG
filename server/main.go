package main

import (
	"OpenRPG/src/backend/classes"
	"fmt"
	"github.com/goccy/go-json"
)

func main() {
	player := classes.Lucky
	classes.CheckValidClass()
	player.RaceAbilities()
	player.CalcLifePoints()

	b, err := json.MarshalIndent(player, "", "     ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}
