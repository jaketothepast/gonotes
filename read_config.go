package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type Config struct {
	NoteDirectory string
}

func WriteConfig() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Directory for notes (blank for ~/.gonotes/notes): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// Blank input
	// TODO make this write config a CLI option
	if input == "\n" {
		fmt.Println("Chosen directory: ~/.gonotes/notes")
		input = path.Join(os.Getenv("HOME"), ".gonotes/notes")
	} else {
		fmt.Println("Chosen directory: ", input)
	}

	// TODO check for existence of notes directory before writing config

	c := Config{NoteDirectory: input}
	conf, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(conf))
	err2 := ioutil.WriteFile(path.Join(os.Getenv("HOME"), ".gonotes/config.json"), conf, 0644)
	if err2 != nil {
		panic(err2)
	}
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
			} else {
				WriteConfig()
			}
		}
	}

	return true
}
