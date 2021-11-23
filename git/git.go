package git

import (
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


func Branch() ([]string, error) {
	out, err := exec.Command("git", "branch").Output()

	if err != nil {
		return nil, err
	}


	output := strings.ReplaceAll(string(out), "*", "")
	output = strings.ReplaceAll(output, " ", "")

	branchesList := filter(strings.Split(output, "\n"), func(item string, _ int) bool {
		return len(strings.TrimSpace(item)) > 0
	})

	return branchesList, nil
}

func Checkout(branch string, createBranch bool) bool {
	if branch == "" {
		return false
	}

	if createBranch {
		if err := exec.Command("git", "checkout", "-b", branch).Run(); err != nil {
			return false
		}

		return true
	}

	if err := exec.Command("git", "checkout", branch).Run(); err != nil {
		return false
	}

	return true
}



func filter(arr []string, predicate func (item string, idx int) bool ) []string {
	var filteredArr []string

	for i, s := range arr {
		if predicate(s, i) {
			filteredArr = append(filteredArr, s)
		}
	}

	return filteredArr
}