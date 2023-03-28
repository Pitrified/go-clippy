package main

import (
	"clippy/gui"
	"clippy/utils"
)

func main() {
	lg := utils.GetStdLogger("")
	lg.Println("Going.")

	theController := gui.NewController()
	theController.Run()
}
