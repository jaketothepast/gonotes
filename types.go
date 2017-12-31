package main

import (
	"time"
	"fmt"
)

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
	lastModified time.Time
	createdAt time.Time

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

/*
 * Set the text and the lastModified flags of a Note object
 */
func (n *Note) SetText(text string) {
	n.text = text
	n.lastModified = time.Now()
}

/*
 * Factory function to create a new note
 */
func NewNote(text string, priority int) *Note {
	note := new(Note)
	note.createdAt = time.Now()
	note.priority = priority
	note.SetText(text)
	return note
}

func main() {
	n := NewNote("hello world", 22)
	fmt.Println(n.createdAt)
}

