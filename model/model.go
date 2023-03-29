package model

import (
	"container/list"
	"log"

	"clippy/utils"
)

// --------------------------------------------------------------------------------
//  Define and create the model
// --------------------------------------------------------------------------------

type GuiModel struct {
	regNum     int
	RegContent *list.List

	lg *log.Logger
}

// Create the model.
func NewModel(regNum int) *GuiModel {
	m := &GuiModel{
		lg:         utils.GetStdLogger(""),
		regNum:     regNum,
		RegContent: list.New(),
	}
	return m
}

// --------------------------------------------------------------------------------
//  React to user input: change the model state
// --------------------------------------------------------------------------------

// The user clicked on a register button.
func (m *GuiModel) ClickedReg(whichReg int) {
	m.lg.Println("Clicked", whichReg)

	// the value is in the list, unless the user has clicked on a button that
	// does not have a value yet: if it is in the list, move it to the front
	i := 0
	var e *list.Element
	for e = m.RegContent.Front(); e != nil; e = e.Next() {
		if i == whichReg {
			m.RegContent.MoveToFront(e)
			return
		}
		i++
	}

}

// The user pasted something.
func (m *GuiModel) Pasted(data string) {
	m.lg.Println("Pasted", utils.FmtRegContent(data))

	// if the data is already in the list, move it to the front
	for e := m.RegContent.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == data {
			m.RegContent.MoveToFront(e)
			return
		}
	}

	// otherwise, add it to the front
	m.RegContent.PushFront(data)

	// if the list is too long, remove the last element
	if m.RegContent.Len() > m.regNum {
		m.RegContent.Remove(m.RegContent.Back())
	}

}

// --------------------------------------------------------------------------------
//  Update the model (internal logic)
// --------------------------------------------------------------------------------

// blah it's a minimal example, there is no internal logic to update
