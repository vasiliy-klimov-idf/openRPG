package core

import (
	cl "github.com/vasiliy-klimov-idf/openrpg_core/src/classes"
	"github.com/vasiliy-klimov-idf/openrpg_core/src/player"
	"github.com/vasiliy-klimov-idf/openrpg_core/src/story"
	"runtime"
)

var (
	P = player.Player{}

	SavePath      string
	SelectedRace  string
	SelectedClass cl.Class
	PageNumber    int
)

func Initialization() {
	if runtime.GOOS == "windows" {
		SavePath = "C:\\"
	} else {
		SavePath = "/usr/share/local/"
	}

}

func InitPlayer() {
	P.Nickname = story.PlayerName
	P.Class.Race = SelectedRace
	P.Class = SelectedClass

	P.CreatePlayer(P.Nickname, P.Class, P.Class.Race)
}
