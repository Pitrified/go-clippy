package utils

import (
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// Get a logger with my version of standard flags.
func GetStdLogger(prefix string) *log.Logger {
	lg := log.New(
		os.Stderr,
		prefix,
		log.LstdFlags|
			log.Ltime|
			log.Lshortfile,
	)
	return lg
}

// Format a string for display in a button.
//
// If the string is too long, truncate the middle.
func FmtRegContent(s string) string {
	pad := 30
	if len(s) > pad*2 {
		s = s[:pad] + " ... " + s[len(s)-pad:]
	}
	s = strings.Replace(s, "\n", " ", -1)
	return s
}

// Convert a slice of a more specific type to a slice of fyne.CanvasObject.
//
// Mildly inefficient, but it works.
func ToCO[T fyne.CanvasObject](objects ...T) []fyne.CanvasObject {
	objCO := make([]fyne.CanvasObject, len(objects))
	for i := range objCO {
		objCO[i] = objects[i]
	}
	return objCO
}

// NewHBox creates a new container with the specified objects and using the HBox layout.
// The objects will be placed in the container from left to right.
func NewHBox[T fyne.CanvasObject](objects ...T) *fyne.Container {
	return container.NewHBox(ToCO(objects...)...)
}

// NewVBox creates a new container with the specified objects and using the VBox layout.
// The objects will be stacked in the container from top to bottom.
func NewVBox[T fyne.CanvasObject](objects ...T) *fyne.Container {
	return container.NewVBox(ToCO(objects...)...)
}
