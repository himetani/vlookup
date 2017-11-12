package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("vlookup,-v", ",")
	cli.Run(args)
	expected := fmt.Sprintf("vlookup version")
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}

}
