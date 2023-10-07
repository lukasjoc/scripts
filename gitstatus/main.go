package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	git "github.com/go-git/go-git/v5"
)

const (
	ansiColorBlack   = "\033[30m"
	ansiColorRed     = "\033[31m"
	ansiColorGreen   = "\033[32m"
	ansiColorYellow  = "\033[33m"
	ansiColorBlue    = "\033[34m"
	ansiColorMagenta = "\033[35m"
	ansiColorCyan    = "\033[36m"
	ansiColorWhite   = "\033[37m"
	ansiReset        = "\033[0m"
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

func main() {
	type GitStatus struct {
		Branch   ansiText
		HeadHash ansiText
		Flags    []ansiText
	}
	tmpl, err := template.New("gitstatus").Parse("({{.Branch}}@{{.HeadHash}}{{if gt (len .Flags) 0}};{{range .Flags}}{{.}}{{end}}{{end}})\n")
	if err != nil {
		panic(err)
	}

	path, err := revParseShowToplevel()
	if err != nil {
		panic(err)
	}
	repo, err := git.PlainOpen(path)
	if err != nil {
		if err == git.ErrRepositoryNotExists {
			os.Exit(0)
		}
		panic(err)
	}

	worktree, _ := repo.Worktree()
	status, err := worktree.Status()
	if err != nil {
		panic(err)
	}

	flags := map[git.StatusCode]ansiText{}
	// TODO: flag for if behind remote and how many commits
	for _, stat := range status {
		switch stat.Worktree {
		case git.Untracked:
			flags[git.Untracked] = ansiText{ansiColorBlue, "+"}
		case git.Modified:
			flags[git.Modified] = ansiText{ansiColorYellow, "*"}
		case git.Added:
			flags[git.Added] = ansiText{ansiColorCyan, "+"}
		case git.Deleted:
			flags[git.Deleted] = ansiText{ansiColorGreen, "-"}
		case git.Renamed:
			flags[git.Renamed] = ansiText{ansiColorYellow, "r"}
		}
	}
	flagsFormatted := []ansiText{}
	for _, flag := range flags {
		flagsFormatted = append(flagsFormatted, flag)
	}

	head, err := repo.Head()
	if err != nil {
		panic(err)
	}

	formatted := GitStatus{
		ansiText{ansiColorYellow, head.Name().Short()},
		ansiText{ansiColorGreen, head.Hash().String()[0:8]},
		flagsFormatted,
	}
	if err := tmpl.Execute(os.Stdout, formatted); err != nil {
		panic(err)
	}
}
