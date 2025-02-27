// asciitell prints the first letter of the first argument string in decimal
// for faster access of an ascii character without having to search through `man ascii`.
// Probably should just memorize this by now.
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: [symbol]\n")
		os.Exit(1)
	}
	fmt.Printf("%d\n", os.Args[1][0])
}
