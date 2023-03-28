package utils

import (
	"log"
	"os"
	"strings"
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
