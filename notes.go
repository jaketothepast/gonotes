package main

import (
	"time"
	"path"
	"fmt"
	"strings"
	"strconv"
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

}
