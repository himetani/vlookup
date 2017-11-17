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

func init() {
	flags.BoolVar(&printVersion, "version", false, "print version")
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}

func (c *CLI) Run(args []string) int {
	flags.Usage = func() {
		fmt.Fprintf(c.errStream, "vlookup version %s\n\n", Version)
		fmt.Fprintf(c.errStream, "Usage: %s [value](list file) [table](csv file) index_number\n\n", os.Args[0])
		flags.PrintDefaults()
	}

	flags.SetOutput(c.errStream)

	if err := flags.Parse(args[1:]); err == flag.ErrHelp {
		return 1
	}

	if printVersion {
		fmt.Fprintf(c.outStream, "vlookup version %s\n", Version)
		return 0
	}

	if len(args) != 4 {
		return 1
	}

	// params
	value := args[1]
	table := args[2]
	i, err := strconv.Atoi(args[3])
	if err != nil {
		return 1
	}

	// value
	f, err := os.Open(value)
	if err != nil {
		return 1
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	values := []string{}
	for s.Scan() {
		values = append(values, s.Text())
	}

	// table
	f, err = os.Open(table)
	if err != nil {
		return 1
	}
	defer f.Close()

	s = bufio.NewScanner(f)
	tables := [][]string{}
	for s.Scan() {
		tables = append(tables, strings.Split(s.Text(), ","))
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

	return 0
}
