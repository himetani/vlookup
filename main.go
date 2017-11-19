package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	Version     = "0.0.9"
	Name        = "vlookup"
	AuthName    = "himetani"
	AuthorEmail = "takafumi_t1224@jcom.home.ne.jp"
)

var flags = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

var (
	printVersion bool
)

type CLI struct {
	outStream, errStream io.Writer
}

const (
	ExitCodeOk int = iota
	ExitCodeError
)

func init() {
	flags.BoolVar(&printVersion, "v", false, "print version")
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}

	os.Exit(cli.Run(os.Args))
}

func (c *CLI) Run(args []string) int {
	flags.Usage = func() {
		fmt.Fprintf(c.errStream, "Usage: %s [value](list file) [table](csv file) index_number\n", os.Args[0])
		flags.PrintDefaults()
	}

	flags.SetOutput(c.errStream)

	if err := flags.Parse(args[1:]); err == flag.ErrHelp {
		return ExitCodeOk
	}

	if printVersion {
		fmt.Fprintf(c.outStream, "vlookup version %s\n", Version)
		return ExitCodeOk
	}

	if len(args) != 4 {
		fmt.Fprintf(c.errStream, "Arguments number is invalid\n")
		return ExitCodeError
	}

	// params
	value := args[1]
	table := args[2]
	i, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Fprintf(c.errStream, "index_number should be integer value\n")
		return ExitCodeError
	}

	// value
	f, err := os.Open(value)
	if err != nil {
		fmt.Fprintf(c.errStream, "value file can't open\n")
		fmt.Fprintf(c.errStream, "%s\n", err.Error())
		return ExitCodeError
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	values := []string{}
	for s.Scan() {
		values = append(values, s.Text())
	}

	// table

	// min column
	min := 0

	f, err = os.Open(table)
	if err != nil {
		fmt.Fprintf(c.errStream, "table file can't open\n")
		fmt.Fprintf(c.errStream, "%s\n", err.Error())
		return ExitCodeError
	}
	defer f.Close()
	s = bufio.NewScanner(f)
	tables := [][]string{}
	for s.Scan() {
		row := strings.Split(s.Text(), ",")
		tables = append(tables, row)
		if min == 0 {
			min = len(row)
		}

		if len(row) < min {
			min = len(row)
		}
	}

	if min < i {
		fmt.Fprintf(c.errStream, "index_number is invalid\n")
		return ExitCodeError
	}

	buf := new(bytes.Buffer)
	for _, v := range values {
		unmatch := true
		for _, row := range tables {
			if row[0] == v {
				buf.WriteString(fmt.Sprintf("%s,%s\n", v, row[i-1]))
				unmatch = false
				break
			}
		}
		if unmatch {
			buf.WriteString(fmt.Sprintf("%s,nil\n", v))
		}
	}

	fmt.Fprintf(c.outStream, buf.String())

	return ExitCodeOk
}
