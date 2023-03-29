package gui

import (
	"container/list"
	"context"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.design/x/clipboard"

	"clippy/utils"
)

type registers struct {
	a         *guiApp
	container *fyne.Container

	regNum    int
	regBtn    []*widget.Button
	lastPaste string

	lg *log.Logger
}

// Create a new registers pack.
func newRegisters(a *guiApp) *registers {
	r := &registers{
		a:         a,
		lg:        utils.GetStdLogger(""),
		regNum:    5,
		lastPaste: "",
	}

	// create the buttons
	r.regBtn = make([]*widget.Button, r.regNum)
	for i := range r.regBtn {
		i := i // save local copy
		r.regBtn[i] = widget.NewButton("", func() { r.regCB(i) })
	}

	// assemble them
	r.container = utils.NewVBox(r.regBtn...)

	return r
}

// Listen to clipboard changes.
func (r *registers) clipboardWatch() {

	// Init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	// Watch returns a channel that will receive clipboard updates.
	ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
	for data := range ch {
		strData := string(data)
		if strData == r.lastPaste {
			continue
		}
		r.lg.Println("Clipboard changed:", utils.FmtRegContent(strData))
		r.lastPaste = strData
		r.a.c.pasted(strData)
	}
}

// -------------------------------------------------------------------
//  Reactions to user input:
//  callbacks to communicate with the Controller
// -------------------------------------------------------------------

// Clicked one of the register buttons.
func (r *registers) regCB(whichReg int) {
	r.lg.Printf("Clicked %d\n", whichReg)
	// clicking the first register does nothing
	if whichReg == 0 {
		return
	}
	r.a.c.clickedReg(whichReg)
}

// -------------------------------------------------------------------
//  Update the app UI:
//  new state received from the controller
// -------------------------------------------------------------------

// Update the registers.
func (r *registers) updateRegContent(regContent *list.List) {
	i := 0
	for e := regContent.Front(); e != nil; e = e.Next() {
		newContent := e.Value.(string)
		fmtContent := utils.FmtRegContent(newContent)
		r.regBtn[i].SetText(fmtContent)
		i++
	}

	// might be nil if the user has not pasted anything yet.
	maybeFront := regContent.Front()
	if maybeFront == nil {
		return
	}
	newCont := maybeFront.Value.(string)
	// might have race condition here who knows
	// mark the last paste so we don't get another paste event
	r.lastPaste = newCont
	clipboard.Write(clipboard.FmtText, []byte(newCont))

}
