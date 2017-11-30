package main

import (
	"time"
	"path"
	"fmt"
	"strings"
)

// Append command line note to filename with date
func AppendToDateFile(configPath string, note string) {
	config := ReadConfig(configPath)
	currentDay, currentMonth, currentYear := time.Now().Date()

	// TODO get the current date and use to create/read file
	// under the config directory
	strs := []string{string(currentDay), currentMonth.String(), string(currentYear)}

	notePath := path.Join(config.NoteDirectory, strings.Join(strs, "-"))
	fmt.Println(notePath)
}
