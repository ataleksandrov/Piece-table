# Piece-table
 Golang FMI 2k18 Homework 2

## Overview
Implementation of core functionality of a text editor using piece table - [Golang FMI 2k18 Homework 2](https://fmi.golang.bg/tasks/2)

## Usage

`func NewEditor(s string) Editor` returns a `TextManager` that implements `Editor` interface.
```
type Editor interface {
	// Insert text starting from given position.
	Insert(position uint, text string) Editor

	// Delete length items from offset.
	Delete(offset, length uint) Editor

	// Undo reverts latest change.
	Undo() Editor

	// Redo re-applies latest undone change.
	Redo() Editor

	// String returns complete representation of what a file looks
	// like after all manipulations.
	String() string
}
```