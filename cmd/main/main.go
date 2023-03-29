package main

import (
	"clippy/gui"
	"clippy/utils"
	"flag"
)

func main() {
	lg := utils.GetStdLogger("")
	lg.Println("Going.")

	regNum := flag.Int("r", 5, "Number of registries.")
	flag.Parse()

	theController := gui.NewController(*regNum)
	theController.Run()
}
