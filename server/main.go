package main

import (
	"github.com/vasiliy-klimov-idf/openrpg_core/src/classes"
	"github.com/vasiliy-klimov-idf/openrpg_core/src/player"
)

var p = player.Player{}

func init() {
	classes.CheckValidClass()
}

func main() {

	p.CreatePlayer("red23", classes.Mage)
	p.GetPlayerInfo()
}
