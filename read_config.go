package main

import (
	"fmt"
	"os"
)

type Config struct {
	NoteDirectory string
}

func WriteConfig() {

}

func CheckConfig(configPath string) bool {
	if _, err := os.Stat(configPath); err != nil {
		fmt.Println("Could not open file " + configPath)
		fmt.Println(err)

		if configPath != "~/.gonotes/config.json" {
			return false
		} else {
			os.Mkdir("~/.gonotes", 0600)
			os.Create("~/.gonotes/config.json")
		}
	}

	return true
}
