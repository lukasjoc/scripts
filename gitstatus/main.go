package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"

	git "github.com/go-git/go-git/v5"
)

const (
	ansiColorRed    = "\033[31m"
	ansiColorGreen  = "\033[32m"
	ansiColorYellow = "\033[33m"
	ansiColorCyan   = "\033[36m"
	ansiColorWhite  = "\033[37m"
	ansiReset       = "\033[0m"
)

const (
	StatusCodeUntracked = "??"
	StatusCodeModified  = "M"
	StatusCodeDeleted   = "D"
	StatusCodeRenamed   = "R"
)

type ansiText struct {
	Color string
	Text  string
}

func (ansi ansiText) String() string {
	return fmt.Sprintf("%s%s%s", ansi.Color, ansi.Text, ansiReset)
}

// FIXME: go-git issue https://github.com/go-git/go-git/issues/74
func revParseShowToplevel() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", nil
	}
	return strings.Trim(string(out), "\n"), nil
}

// TODO: i need good package with git command wrappers thats not as heavy as go-git
// and relies on the git binary itself.
// NOTE: using this is waaay faster then worktree status from go-git. This is disappointing :(
func status() (string, error) {
	out, err := exec.Command("git", "status", "-s", "-u", "-M", "--porcelain").Output()
	if err != nil {
		return "", nil
	}
	return strings.Trim(string(out), "\n"), nil
}

func main() {
	path, err := revParseShowToplevel()
	if err != nil {
		os.Exit(1)
	}
	repo, err := git.PlainOpen(path)
	if err != nil {
		if err == git.ErrRepositoryNotExists {
			os.Exit(0)
		}
		os.Exit(1)
	}

	status, err := status()
	if err != nil {
		os.Exit(1)
	}
	statuses := []string{}
	for _, line := range strings.Split(status, "\n") {
		l := strings.Split(strings.Trim(line, " "), " ")
		stat := l[0]
		if !slices.Contains(statuses, stat) {
			statuses = append(statuses, stat)
		}
	}

	flags := []ansiText{}
	// TODO: flag for if behind remote and how many commits
	for _, stat := range statuses {
		switch stat {
		case StatusCodeUntracked:
			flags = append(flags, ansiText{ansiColorYellow, "+"})
		case StatusCodeModified:
			flags = append(flags, ansiText{ansiColorCyan, "*"})
		case StatusCodeDeleted:
			flags = append(flags, ansiText{ansiColorRed, "-"})
		case StatusCodeRenamed:
			flags = append(flags, ansiText{ansiColorRed, "r"})
		}
	}

	// TODO: should replace this with git command to
	head, err := repo.Head()
	if err != nil {
		os.Exit(1)
	}

	fmt.Print(ansiText{ansiColorYellow, head.Name().Short()})
	fmt.Print("::")
	fmt.Print(ansiText{ansiColorGreen, head.Hash().String()[0:8]})
	if len(flags) > 0 {
		fmt.Print(flags)
	}
}
