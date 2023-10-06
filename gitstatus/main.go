package main

import (
	"fmt"
	"os"
	"text/template"
)

const (
	AnsiColorBlack   = "\033[30m"
	AnsiColorRed     = "\033[31m"
	AnsiColorGreen   = "\033[32m"
	AnsiColorYellow  = "\033[33m"
	AnsiColorBlue    = "\033[34m"
	AnsiColorMagenta = "\033[35m"
	AnsiColorCyan    = "\033[36m"
	AnsiColorWhite   = "\033[37m"
	AnsiReset        = "\033[0m"
)

type AnsiText struct {
	color string
	Text  string
}

func (ansi AnsiText) String() string {
	return fmt.Sprintf("%s%s%s", ansi.color, ansi.Text, AnsiReset)
}

func main() {
	type Status struct {
		Branch  AnsiText
		Flags   []AnsiText
		HeadSha AnsiText
	}
    tmpl, err := template.New("gitstatus").Parse("({{.Branch}}@{{.HeadSha}}; {{range .Flags}}{{.}}{{end}})\n")
	if err != nil {
		panic(err)
	}

	status := Status{
		AnsiText{AnsiColorYellow, "main"},
		[]AnsiText{
			{AnsiColorBlue, "*"},
			{AnsiColorYellow, "+"},
			{AnsiColorMagenta, "^"},
		},
		AnsiText{AnsiColorGreen, "6feab33"},
	}
	if err := tmpl.Execute(os.Stdout, status); err != nil {
		panic(err)
	}
}
