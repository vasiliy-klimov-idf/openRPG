package basic

import (
	"fmt"
	"time"
)

const progressPercent = 100

func progressBarLineGenerator() string {
	var line = "|"
	for i := 0; i < progressPercent; i++ {
		fmt.Print(line)
		time.Sleep(100 * time.Millisecond)
	}
	return line
}

func InterfacePanel() {
	playerName := "RED"
	enemyName := "Bear"
	fmt.Println(
		"_________________________________________________"+"\n"+
			"|\t", playerName, "\t\t", "|\t", enemyName, "\t\t|"+"\n"+
			"-------------------------------------------------"+"\n"+
			"| HP:", playerName, "\t\t", "| HP:", enemyName, "\t\t|"+"\n"+
			"| MP:", playerName, "\t\t", "| MP:", enemyName, "\t\t|"+"\n"+
			"| SP:", playerName, "\t\t", "| SP:", enemyName, "\t\t|"+"\n"+
			"| PD:", playerName, "\t\t", "| PD:", enemyName, "\t\t|"+"\n"+
			"| MD:", playerName, "\t\t", "| MD:", enemyName, "\t\t|"+"\n"+
			"-------------------------------------------------")
}
