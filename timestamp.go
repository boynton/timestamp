package main

import "fmt"
import "time"

func main() {
	t := time.Now().UTC()
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d.%03dZ\n",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond()/1000000)
}
