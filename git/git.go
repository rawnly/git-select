package git

import (
	"github.com/rawnly/git-select/array"
	"github.com/rawnly/git-select/term"
	"github.com/ssoroka/slice"
	"os/exec"
	"strings"
)

func GetCurrentBranch() string {
	var currentBranch string
	b, err := exec.Command("git", "branch").Output()

	if err != nil {
		return currentBranch
	}

	output := strings.ReplaceAll(string(b), " ", "")

	for _, b := range strings.Split(output, "\n") {
		if strings.Contains(b, "*") {
			currentBranch = b
		}
	}

	return strings.ReplaceAll(currentBranch, "*", "")
}

func Commits() (map[string]string, error) {
	out, err := exec.Command("git", "log", "--all", "--oneline").Output()

	if err != nil {
		return nil, err
	}

	output := strings.TrimSpace(string(out))

	commits := array.Filter(strings.Split(output, "\n"), func(commit string, idx int) bool {
		return len(commit) > 0
	})

	commitMap := array.Reduce[string, map[string]string](commits, make(map[string]string), func(i int, acc map[string]string, commit string) map[string]string {
		row := strings.Split(commit, " ")

		id, message := slice.Shift(row)

		acc[id] = strings.Join(message, " ")

		return acc
	})

	return commitMap, nil
}

func Branch() ([]string, error) {
	out, err := term.RunCommand("git", "branch")

	if err != nil {
		return nil, err
	}

	output := strings.ReplaceAll(string(out), "*", "")
	output = strings.ReplaceAll(output, " ", "")

	branchesList := array.Filter(strings.Split(output, "\n"), func(item string, _ int) bool {
		return len(strings.TrimSpace(item)) > 0
	})

	return branchesList, nil
}

func Checkout(branch string, createBranch bool) bool {
	if branch == "" {
		return false
	}

	if createBranch {
		if err := term.RunOSCommand("git", "checkout", "-b", branch); err != nil {
			return false
		}

		return true
	}

	if err := term.RunOSCommand("git", "checkout", branch); err != nil {
		return false
	}

	return true
}
