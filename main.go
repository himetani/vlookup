package main

import (
	"flag"
	"fmt"
	"io"
	"os"
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
	flags.BoolVar(&printVersion, "v", false, "print version")

}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}

func (c *CLI) Run(args []string) int {
	flags.Usage = func() {
		fmt.Fprintf(c.errStream, "vlookup version %s\n\n", Version)
		fmt.Fprintf(c.errStream, "Usage: %s [value](list file) [table](csv file)\n\n", os.Args[0])
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

	return 0
}
