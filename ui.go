package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ttacon/chalk"
)

type Ui struct {
	Writer      io.Writer
	ErrorWriter io.Writer
}

func NewUi() *Ui {
	return &Ui{
		Writer: os.Stdout,
		// FIXME os.Stderr doesn't print anything
		ErrorWriter: os.Stdout,
	}
}

func output(w io.Writer, msg string) {
	fmt.Fprint(w, msg)
	fmt.Fprint(w, "\n")
}

func (self *Ui) Error(format string, a ...interface{}) {
	w := self.Writer
	if self.ErrorWriter != nil {
		w = self.ErrorWriter
	}
	message := chalk.Red.Color(fmt.Sprintf(format, a...))
	output(w, message)
}

func (self *Ui) Info(format string, a ...interface{}) {
	message := chalk.Cyan.Color(fmt.Sprintf(format, a...))
	output(self.Writer, message)
}
