package main

/*
This application will take notes from the command line, or from a file, and add
those notes to a note file
*/

import (
	"flag"
	"fmt"
	"log"
)

var noteFileName string
var note string
var configFilePath string

type noteFile struct {
	filename string
	notes    []byte
}

func init() {
	flag.StringVar(&noteFileName, "noteFileName", "example_file", "Filepath to notefile")
	flag.StringVar(&note, "note", "example_note", "Note to insert")
	flag.StringVar(&configFilePath, "configFilePath", "~/.gonotes/config.json", "Filepath to config")
}

func main() {
	flag.Parse()
	if note == "example_note" && noteFileName == "example_file" {
		log.Fatal("Please specify a note file, or a note to save")
	}

	fmt.Println(note)
	configPresent := CheckConfig(configFilePath)
	if configPresent {
		log.Print("Config is present, ready for note-taking!")
	}
}
