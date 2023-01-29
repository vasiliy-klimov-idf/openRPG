package backend

import (
	dis "github.com/fstanis/screenresolution"
)

func GetDisplaySize() string {
	return dis.GetPrimary().String()
}

func GetDisplayWidth() int {
	return dis.GetPrimary().Width
}
func GetDisplayHeight() int {
	return dis.GetPrimary().Height
}
