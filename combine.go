// combine combines lines of stdin with the window provided
// e.g. echo -e "foo\nbar\nbaz" | combine 2 would combine to
//
//	foo bar
//	baz
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Line is a line in the output of this program
type Line string

// combine combines the output of a scanner using the window `w` provided.
// w > len(lines) combines into a single line
// w == 0 doesnt combine at all
func combine(scanner *bufio.Scanner, w int, sep string) (lines []Line) {
	buf := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if w > 1 {
			buf = append(buf, line)
			if len(buf) == w {
				lines = append(lines, Line(strings.Join(buf, sep)))
				buf = nil
			}
		} else {
			lines = append(lines, Line(line))
		}
	}
	if buf != nil {
		lines = append(lines, Line(strings.Join(buf, sep)))
		buf = nil
	}
	return lines
}

func main() {
	warg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for _, line := range combine(scanner, warg, " ") {
		fmt.Println(line)
	}
}
