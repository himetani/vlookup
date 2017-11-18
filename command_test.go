package main

import (
	"bytes"
	"strings"
	"testing"
)

const (
	stdout int = iota
	stderr
)

func TestCommands(t *testing.T) {
	commadtests := []struct {
		name     string
		args     string
		expected string
		std      int
		status   int
	}{
		{"Success", "vlookup,example/value.csv,example/table.csv,1", "111,111\n222,nil\n333,333", stdout, ExitCodeOk},
		{"Argument", "vlookup,example/value.csv,example/table.csv", "Arguments number is invalid", stderr, ExitCodeError},
		{"index_number", "vlookup,example/value.csv,example/table.csv,hoge", "index_number should be integer value", stderr, ExitCodeError},
		{"value", "vlookup,example/hoge.csv,example/table.csv,1", "value file can't open", stderr, ExitCodeError},
		{"table", "vlookup,example/value.csv,example/hoge.csv,1", "table file can't open", stderr, ExitCodeError},
		{"missing table", "vlookup,example/value.csv,example/missing_table.csv,2", "index_number is invalid", stderr, ExitCodeError},
	}

	for _, c := range commadtests {
		outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
		cli := &CLI{outStream: outStream, errStream: errStream}
		args := strings.Split(c.args, ",")
		status := cli.Run(args)

		switch c.std {
		case stdout:
			if !strings.Contains(outStream.String(), c.expected) {
				t.Errorf("Case: %s\n", c.name)
				t.Errorf("expected %q to eq %q", outStream.String(), c.expected)
			}
		case stderr:
			if !strings.Contains(errStream.String(), c.expected) {
				t.Errorf("Case: %s\n", c.name)
				t.Errorf("expected %q to eq %q", outStream.String(), c.expected)
			}
		}

		if status != c.status {
			t.Errorf("Case: %s\n", c.name)
			t.Errorf("expected %d to eq %d", status, c.status)
		}
	}
}
