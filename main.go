package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/rawnly/git-select/git"
	"os"
)


func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		branchToCheckout := args[0]

		if args[0] == "-b" {
			branchToCheckout = args[1]

			if git.Checkout(branchToCheckout, true) {
				fmt.Printf("Switched to branch '%s'", branchToCheckout)
				return
			}
		} else if branchToCheckout != "" {
			if git.Checkout(branchToCheckout, false) {
				fmt.Printf("Switched to branch '%s'", branchToCheckout)
				return
			}
		}
	}


	branches, err := git.Branch()

	if err != nil {
		panic(err)
	}


	qs := []*survey.Question{
		{
			Name: "branch",
			Prompt: &survey.Select{
				Message: "Select a branch",
				Options: branches,
				Default: git.GetCurrentBranch(),
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

	git.Checkout(answ.Branch, false)
}


