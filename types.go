package main

// One leaf on the tree, represents a note/task
type Note struct {
	text string
	priority int
	task bool
}

// Hold notes, can have child pages
type NotePage struct {
	title string
	children []NotePage
	filename string
	notes []Note
}

