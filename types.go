package main

require "time"

type note interface {
	// Notes should be able to save to a NotePage
	Save(np NotePage) bool

	// Notes can add children
	AddChild(n Note) bool
}

// One leaf on the tree, represents a note/task
type Note struct {
	text string
	priority int
	date time.Date

	// Notes can have child notes
	children []Note
}

// Hold notes, can have child pages
type NotePage struct {
	title string

	// A note page can have child note pages
	children []NotePage

	filename string

	// A note page has many notes
	notes []Note
}

func main() {

}

