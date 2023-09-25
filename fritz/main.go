package main

import (
	"fmt"
	"github.com/lukasjoc/scripts/fritz/internal"
	"os"
)

func printHelp() {
    fmt.Fprintf(os.Stderr, "\n%s [command]\n", os.Args[0])
    fmt.Fprintf(os.Stderr, "%-10s %s\n", "info", "Print info about the box and its configuration")
    fmt.Fprintf(os.Stderr, "%-10s %s\n", "reconnect", "Quickly Disconnect and Reconnect again")
    fmt.Fprintf(os.Stderr, "%-10s %s\n", "reboot", "Quickly reboot the box")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
	}

	fritz, err := internal.NewFritz()
	if err != nil {
        panic(err)
	}
	switch os.Args[1] {
	case "reboot":
		if err := fritz.Reboot(); err != nil {
			panic(err)
		}
	case "reconnect":
		if err := fritz.Reconnect(); err != nil {
			panic(err)
		}
	case "info":
		if err := fritz.Info(); err != nil {
			panic(err)
		}
	default:
		printHelp()
	}
}
