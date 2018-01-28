package main

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

type Printer struct {
	errout io.Writer
}

func NewPrinter(stderr io.Writer) *Printer {
	color.Output = stderr
	return &Printer{
		errout: stderr,
	}
}

func (p *Printer) Println(args ...interface{}) {
	fmt.Fprintln(p.errout, args...)
}

func (p *Printer) Printf(format string, args ...interface{}) {
	fmt.Fprintf(p.errout, format, args...)
}
