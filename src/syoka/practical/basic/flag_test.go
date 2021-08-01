package basic

import (
	"flag"
	"testing"
)

//go test -run TestFlagSimple -v -args "java"
func TestFlagSimple(t *testing.T) {
	if !flag.Parsed() {
		flag.Parse()
	}

	argList := flag.Args()
	for _, arg := range argList {
		if arg == "java" {
			t.Log("Running in java.")
		} else {
			t.Log("Running in go.")
		}
	}
}
