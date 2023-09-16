package main

import (
	"fmt"
	"github.com/lukasjoc/scripts/fritz/internal"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "You must specify one of the following options: reboot, reconnect")
		os.Exit(1)
	}

	fritz, err := internal.NewFritz()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	cmd := os.Args[1]
	switch cmd {
	case "reboot":
		fritz.Reboot()
	case "reconnect":
		fritz.Reconnect()
	default:
        if err := fritz.Connect(); err != nil {
            panic(err)
        }
        // os.Exit(1)
		// fmt.Fprintln(os.Stderr, "Invalid command. Use 'reboot' or 'reconnect'.")
		// os.Exit(1)
	}
}
