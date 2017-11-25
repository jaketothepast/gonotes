package main

import (
	"fmt"
	"os"
	"path"
)

type Config struct {
	NoteDirectory string
}

func WriteConfig() {

}

func CheckConfigExists(configPath string) bool {
	if _, err := os.Stat(configPath); err != nil {
		fmt.Println("Could not open file " + configPath)
		fmt.Println(err)

		if configPath != "~/.gonotes/config.json" {
			return false
		} else {
			configPath = path.Join(os.Getenv("HOME"), ".gonotes/config.json")
			err = os.Mkdir(path.Dir(configPath), 0700)
			if err != nil {
				fmt.Println(err)
			}
			_, err = os.Create(configPath)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return true
}
