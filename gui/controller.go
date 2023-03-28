package gui

import (
	"clippy/model"
	"log"

	"clippy/utils"
)

type guiController struct {
	a *guiApp
	m *model.GuiModel

	lg *log.Logger
}

// Create a new controller, linked to the view and the model.
func NewController() *guiController {

	// create the controller
	c := &guiController{
		lg: utils.GetStdLogger(""),
	}

	// create the view
	c.a = newApp(c)

	// initialize the model
	c.m = model.NewModel()

	// update all the moving parts to match the current state:
	// the model has valid default values,
	// the view has only placeholders
	c.initAll()

	return c
}

// Run the app.
func (c *guiController) Run() {
	// run the app (will block)
	c.a.runApp()
}

// Init all elements.
func (c *guiController) initAll() {
	c.updateRegContent()
}

// --------------------------------------------------------------------------------
//  Reacts to events from UI (the view calls these funcs from the callbacks):
//  change the state of the model, then update the view.
// --------------------------------------------------------------------------------

// The user clicked on a register button.
func (c *guiController) clickedReg(whichReg int) {
	c.lg.Println("Clicked", whichReg)
	c.m.ClickedReg(whichReg)
	c.updateRegContent()
}

// The user copied something to the clipboard.
func (c *guiController) pasted(data string) {
	c.lg.Println("Pasted", utils.FmtRegContent(data))
	c.m.Pasted(data)
	c.updateRegContent()
}

// --------------------------------------------------------------------------------
//  The model has changed:
//  the controller knows which view elements must be updated.
// --------------------------------------------------------------------------------

// Update the content of the registers.
func (c *guiController) updateRegContent() {
	c.a.updateRegContent(c.m.RegContent)
}
