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
	Color string
	Text  string
}

func (ansi AnsiText) String() string {
	return fmt.Sprintf("%s%s%s", ansi.Color, ansi.Text, AnsiReset)
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
		Branch   AnsiText
		HeadHash AnsiText
		Flags    []AnsiText
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
		} else {
			panic(err)
		}
	}

	worktree, _ := repo.Worktree()
	status, err := worktree.Status()
	if err != nil {
		panic(err)
	}

	flags := map[git.StatusCode]AnsiText{}
	// TODO: flag for if behind remote and how many commits
	for _, stat := range status {
		switch stat.Worktree {
		case git.Untracked:
			flags[git.Untracked] = AnsiText{AnsiColorBlue, "+"}
			break
		case git.Modified:
			flags[git.Modified] = AnsiText{AnsiColorYellow, "*"}
			break
		case git.Added:
			flags[git.Added] = AnsiText{AnsiColorCyan, "+"}
			break
		case git.Deleted:
			flags[git.Deleted] = AnsiText{AnsiColorGreen, "-"}
			break
		case git.Renamed:
			flags[git.Renamed] = AnsiText{AnsiColorYellow, "r"}
			break
		}
	}
	flagsFormatted := []AnsiText{}
	for _, flag := range flags {
		flagsFormatted = append(flagsFormatted, flag)
	}

	head, err := repo.Head()
	if err != nil {
		panic(err)
	}

	formatted := GitStatus{
		AnsiText{AnsiColorYellow, head.Name().Short()},
		AnsiText{AnsiColorGreen, head.Hash().String()[0:8]},
		flagsFormatted,
	}
	if err := tmpl.Execute(os.Stdout, formatted); err != nil {
		panic(err)
	}
}
