package core

import (
	cl "github.com/vasiliy-klimov-idf/openrpg_core/src/classes"
	"runtime"
)

var (
	SavePath      string
	SelectedRace  string
	SelectedClass cl.Class
)

func Initialization() {
	if runtime.GOOS == "windows" {
		SavePath = "C:\\"
	} else {

	}
}
