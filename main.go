package main

/*
This application will take notes from the command line, or from a file, and add
those notes to a note file
*/
import (
	"flag"
	"log"
	"os"
	"path"
)

var noteFileName string
var note string
var configFilePath string
var writeConfig bool
var printToday bool
var printDate string

type noteFile struct {
	filename string
	notes    []byte
}

func init() {
	defaultConfig := path.Join(os.Getenv("HOME"), ".gonotes/config.json")

	flag.StringVar(&noteFileName, "noteFileName", "", "Filepath to notefile")
	flag.StringVar(&note, "note", "", "Note to insert")
	flag.StringVar(&configFilePath, "configFilePath", defaultConfig, "Filepath to config")
	flag.BoolVar(&writeConfig, "writeConfig", false, "Write out a config file for gonotes")
	flag.BoolVar(&printToday, "printToday", false, "Print notes taken today")
	flag.StringVar(&printDate, "printDate", "", "Date to print")
}

func main() {
	flag.Parse()

	if writeConfig == true {
		WriteConfig()
	}

	if note == "" && noteFileName == "" {
		log.Fatal("Please specify a note file, or a note to save")
	}

	// We should make this just one option.
	CheckConfigExists(configFilePath)

	if note != "" {
		AppendToDateFile(configFilePath, note)
	}

	if date != "" {
		PrintDateFile(configPath, printDate)
	}
}
