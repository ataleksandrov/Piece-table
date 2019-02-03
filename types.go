package main

type Piece struct {
	origin bool
	offset int
	length int
}

type State struct {
	pieces []Piece
}

type TextManager struct {
	originBuffer []byte
	addBuffer    []byte
	pieceTable   []State
	redoStack    []State
}
