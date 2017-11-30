package main

import (
	"time"
	"path"
	"fmt"
	"strings"
	"strconv"
	"os"
)

// Append command line note to filename with date
func AppendToDateFile(configPath string, note string) {
	config := ReadConfig(configPath)
	currentDay, currentMonth, currentYear := time.Now().Date()

	// TODO get the current date and use to create/read file
	// under the config directory
	strs := []string{strconv.Itoa(currentDay), currentMonth.String(), strconv.Itoa(currentYear)}

	notePath := path.Join(config.NoteDirectory, strings.Join(strs, "-"))
	fmt.Println(notePath)

	// Open file with append and wr flags set, 0600
	f, err := os.OpenFile(notePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err := f.WriteString(note); err != nil {
		panic(err)
	}
}
