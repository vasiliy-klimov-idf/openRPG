package core

import (
	"github.com/vasiliy-klimov-idf/openrpg_core/src/player"
	"io/ioutil"
	"log"
)

func SaveGame(p player.Player) {
	err := ioutil.WriteFile(SavePath+"info.txt", p.GetPlayerInfo(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
