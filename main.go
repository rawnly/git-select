package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"
	"os/exec"
	"strings"
)

func gitBranch() ([]string, error) {
	b, err := exec.Command("git", "branch").Output()

	if err != nil {
		return nil, err
	}

	output := string(b)

	output = strings.ReplaceAll(output, "*", "")
	output = strings.ReplaceAll(output, " ", "")

	branchesList := filter(strings.Split(output, "\n"), func(item string, _ int) bool {
		return len(strings.TrimSpace(item)) > 0
	})

	return branchesList, nil
}

func getCurrentBranch() string {
	var currentbranch string
	b, err := exec.Command("git", "branch").Output()

	if err != nil {
		return currentbranch
	}

	output := string(b)
	output = strings.ReplaceAll(output, " ", "")

	for _, b := range strings.Split(output, "\n") {
		if strings.Contains(b, "*") {
			currentbranch = b
		}
	}

	return strings.ReplaceAll(currentbranch, "*", "")
}

func gitCheckout(branch string, createBranch bool) bool {
	if branch == "" {
		return false
	}

	if createBranch {
		err := exec.Command("git", "checkout", "-b", branch).Run()

		if err != nil {
			return false
		}
	} else {
		err := exec.Command("git", "checkout", branch).Run()

		if err != nil {
			return false
		}
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

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		branchToCheckout := args[0]

		if args[0] == "-b" {
			branchToCheckout = args[1]

			if gitCheckout(branchToCheckout, true) {
				fmt.Printf("Switched to branch '%s'", branchToCheckout)
				return
			}
		} else if branchToCheckout != "" {
			if gitCheckout(branchToCheckout, false) {
				fmt.Printf("Switched to branch '%s'", branchToCheckout)
				return
			}
		}
	}


	branches, err := gitBranch()

	if err != nil {
		panic(err)
	}


	qs := []*survey.Question{
		{
			Name: "branch",
			Prompt: &survey.Select{
				Message: "Select a branch",
				Options: branches,
				Default: getCurrentBranch(),
			},
		},
	}

	answ := struct {
		Branch string `survey:"branch"`
	}{}

	err = survey.Ask(qs, &answ)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	gitCheckout(answ.Branch, false)
}


