package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now().UTC()
	var format string
	switch len(os.Args) {
	case 1:
		format = "%d-%02d-%02dT%02d:%02d:%02d.%03dZ\n"
	case 2:
		if os.Args[1] != "-f" {
			usage()
		}
		format = "%d%02d%02dT%02d%02d%02d%03d"
	default:
		usage()
	}
	fmt.Printf(format, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1000000)
}

func usage() {
	fmt.Printf("usage:\n")
	fmt.Printf("    timestamp        show the current time in RFC3339 format, UTC, with millisecond resolution, and trailing newline\n")
	fmt.Printf("    timestamp -f     show the current time in a filename friendly format, with no newline\n")
	os.Exit(1)
}
