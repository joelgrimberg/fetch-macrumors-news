package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	exitCode := m.Run()
	if exitCode != 0 {
		os.Exit(exitCode)
	}

	if !strings.Contains(buf.String(), "Hello, world!") {
		fmt.Errorf("main() did not print 'Hello, world!'")
	}
}
