package main

import (
	"bytes"
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	main()

	log.SetOutput(nil) // Reset log output

	expected := "👋 World\n"
	if !bytes.Contains(buf.Bytes(), []byte(expected)) {
		t.Errorf("Expected log to contain %q, but got %q", expected, buf.String())
	}
}
