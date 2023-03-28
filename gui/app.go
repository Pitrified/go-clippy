package gui

import (
	"container/list"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"clippy/utils"
)

type guiApp struct {
	c *guiController

	fyneApp fyne.App
	mainWin fyne.Window

	registers *registers

	lg *log.Logger
}

func newApp(c *guiController) *guiApp {

	// create the fyne app
	fyneApp := app.NewWithID("com.pitrified.clippy")
	mainWin := fyneApp.NewWindow("Clippy")

	// create the clippy app
	a := &guiApp{
		fyneApp: fyneApp,
		mainWin: mainWin,
		c:       c,
		lg:      utils.GetStdLogger(""),
	}

	// add the link for typed runes
	a.mainWin.Canvas().SetOnTypedKey(a.typedKey)

	// create the registers
	a.registers = newRegisters(a)

	// create the UI, using placeholders everywhere
	a.buildUI()

	return a
}

func (a *guiApp) runApp() {
	// start listening to clipboard
	go a.registers.clipboardWatch()

	// run the app
	a.mainWin.Resize(fyne.NewSize(300, 300))
	a.mainWin.Show()
	a.fyneApp.Run()
}

func (a *guiApp) typedKey(ev *fyne.KeyEvent) {
	a.lg.Printf("typedKey  = %+v %T\n", ev, ev)
	switch ev.Name {
	case fyne.KeyEscape:
		a.fyneApp.Quit()
	// case fyne.KeyH:
	// 	a.s.miscHelpCB()
	default:
	}
}

// -------------------------------------------------------------------
//  Build the app UI
// -------------------------------------------------------------------

// Assemble the UI elements.
func (a *guiApp) buildUI() {
	regCont := container.NewVBox(
		layout.NewSpacer(),
		a.registers.container,
		layout.NewSpacer(),
	)
	a.mainWin.SetContent(regCont)
}

// -------------------------------------------------------------------
//  Reactions to user input:
//  callbacks to communicate with the Controller
// -------------------------------------------------------------------

// // Clicked the button requesting a hint.
// func (a *guiApp) hintCB() {
// 	a.c.clickedHint()
// }

// -------------------------------------------------------------------
//  Update the app UI:
//  new state received from the controller
// -------------------------------------------------------------------

// Update the registers.
func (a *guiApp) updateRegContent(regContent *list.List) {
	a.registers.updateRegContent(regContent)
}
