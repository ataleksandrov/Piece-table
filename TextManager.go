package main

func NewEditor(s string) Editor {
	var tm = new(TextManager)
	tm.originBuffer = []byte(s)
	tm.redoStack = make([]State, 0)
	tm.pieceTable = make([]State, 1)
	var st = new(State)
	st.pieces = make([]Piece, 1)
	st.pieces[0] = Piece{true, 0, len(s)}
	tm.pieceTable[0] = *st
	return tm
}

func (tm *TextManager) getCurrentState() State {
	return tm.pieceTable[len(tm.pieceTable)-1]
}

func (tm *TextManager) setNewState(s State) {
	tm.pieceTable = append(tm.pieceTable, s)
}

func (tm *TextManager) Insert(position uint, text string) Editor {
	offset := len(tm.addBuffer) //len(nil) == 0

	tm.redoStack = tm.redoStack[:0]

	if tm.addBuffer == nil {
		tm.addBuffer = []byte(text)
	} else {
		tm.addBuffer = append(tm.addBuffer, []byte(text)...)
	}
	piece := Piece{false, offset, len(text)}

	state := tm.getCurrentState()
	if length := (uint)(len(tm.String())); position >= length {
		state.pieces = append(state.pieces, piece)
	} else if position == 0 {
		state.pieces = append([]Piece{piece}, state.pieces...)
	} else {
		var bytes uint = 0

		for index, p := range state.pieces {
			bytes += (uint)(p.length)
			if bytes > position {
				i := (int)(position - (bytes - (uint)(p.length)))
				p1 := Piece{p.origin, p.offset, i}
				p2 := Piece{p.origin, p.offset + i, p.length - i}
				toAppend := []Piece{p1, piece, p2}
				if index+1 < len(state.pieces) {
					toAppend = append(toAppend, state.pieces[index+1:]...)
				}
				state.pieces = append(state.pieces[:index], toAppend...)
				break
			}
		}
	}
	tm.setNewState(state)
	return tm
}

func (tm *TextManager) Delete(offset, length uint) Editor {
	contentLength := (uint)(len(tm.String()))
	if offset > contentLength {
		return tm
	}
	if offset+length > contentLength {
		length = contentLength - offset
	}

	tm.redoStack = tm.redoStack[:0]

	state := tm.getCurrentState()
	resultState := new(State)
	resultState.pieces = make([]Piece, 0)
	var bytes uint = 0
	flag := false
	for index, piece := range state.pieces {
		bytes += (uint)(piece.length)
		if bytes >= offset {
			i := (int)(offset - (bytes - (uint)(piece.length)))
			var p1, p2 = new(Piece), new(Piece)
			if i != 0 {
				*p1 = Piece{piece.origin, piece.offset, i}
			}
			if length >= (uint)(piece.length-i) {
				offset = bytes
				length -= (uint)(piece.length - i)
				if p1.length != 0 {
					resultState.pieces = append(resultState.pieces, *p1)
				}
				flag = true
				continue
			} else {
				if flag {
					size := (int)((uint)(piece.length) - length)
					*p2 = Piece{piece.origin, piece.offset + (int)(length), size}
				} else {
					size := (int)((uint)(piece.length) - length - offset)
					*p2 = Piece{piece.origin, piece.offset + (int)(offset) + (int)(length), size}
				}
			}
			if p1.length != 0 {
				resultState.pieces = append(resultState.pieces, *p1)
			}
			if p2.length != 0 {
				resultState.pieces = append(resultState.pieces, *p2)
			}
			if index+1 < len(state.pieces) {
				resultState.pieces = append(resultState.pieces, state.pieces[index+1:]...)
			}
			break
		} else {
			offset -= (uint)(piece.length) + 1
			bytes -= (uint)(piece.length)
			resultState.pieces = append(resultState.pieces, piece)
		}
	}
	tm.setNewState(*resultState)
	return tm
}

func (tm *TextManager) Undo() Editor {
	if length := len(tm.pieceTable); length > 1 {
		tm.redoStack = append(tm.redoStack, tm.pieceTable[length-1])
		tm.pieceTable = tm.pieceTable[:length-1]
	}
	return tm
}

func (tm *TextManager) Redo() Editor {
	if length := len(tm.redoStack); length > 0 {
		tm.pieceTable = append(tm.pieceTable, tm.redoStack[length-1])
		tm.redoStack = tm.redoStack[:length-1]
	}
	return tm
}

func (tm *TextManager) String() string {
	i := len(tm.pieceTable) - 1
	res := []byte{}
	for _, p := range tm.pieceTable[i].pieces {
		if p.origin {
			res = append(res, tm.originBuffer[p.offset:p.offset+p.length]...)
		} else {
			res = append(res, tm.addBuffer[p.offset:p.offset+p.length]...)
		}
	}
	return string(res)
}

func main() {
	//println(NewEditor("AB").Insert(2, "CDE").Insert(5, "FG").Delete(1, 5).String())
}
