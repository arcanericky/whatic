package main

import (
	"os"
	"testing"
)

func TestWhatIC(t *testing.T) {
	// Exercise all examples in case tests are run on
	// different platforms
	examples("linux")
	examples("darwin")
	examples("windows")
	help()
	arguments("", "", os.Args)
	main()

	singleDelimiter := "|"
	os.Setenv(envWhatIC, singleDelimiter)
	o, c := delimiters()
	if o != singleDelimiter || c != singleDelimiter {
		t.Error("Single delimiter not set properly")
	}

	doubleDelimiter := "[]"
	os.Setenv(envWhatIC, doubleDelimiter)
	o, c = delimiters()
	if o != string(doubleDelimiter[0]) || c != string(doubleDelimiter[1]) {
		t.Error("Double delimiter not set properly")
	}

	savedArgs := os.Args
	os.Args = []string{"arg0"}
	main()
	os.Args = savedArgs
}
