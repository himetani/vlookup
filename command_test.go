package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCommands(t *testing.T) {
	commadtests := []struct {
		args     string
		expected string
	}{
		{"vlookup,example/value.csv,example/table.csv,1", "111,111\n222,nil\n333,333"},
	}

	for _, c := range commadtests {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{outStream: outStream, errStream: errStream}
		args := strings.Split(c.args, ",")
		cli.Run(args)
		if !strings.Contains(outStream.String(), c.expected) {
			t.Errorf("expected %q to eq %q", outStream.String(), c.expected)
		}
	}
}
