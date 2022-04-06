package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/alecthomas/kong"
	"github.com/mgutz/ansi"
	"github.com/rawnly/git-select/git"
	"github.com/ssoroka/slice"
	"strings"
)

var cli struct {
	Branch  string `flag:"" name:"branch" help:"Branch to checkout" optional:"" short:"b"`
	Commits bool   `help:"Select a commit to checkout" optional:""`
	Version bool   `help:"Shows cli version" flag:"" short:"v"`

	Target string `arg:"" name:"checkout-target" optional:"" help:"Branch/Commit to checkout (an alias for -b)"`
}

var version string = "development"

func main() {
	_ = kong.Parse(&cli)
	highlight := ansi.ColorFunc("black:yellow")

	if cli.Version {
		fmt.Println(fmt.Sprintf("Version: %s", version))
		return
	}

	var checkoutTarget string

	if cli.Target != "" {
		checkoutTarget = cli.Target
	}

	if cli.Branch != "" {
		checkoutTarget = cli.Branch
	}

	if checkoutTarget != "" {
		if git.Checkout(checkoutTarget, true) {
			fmt.Printf("Switched to branch '%s'", cli.Target)
			return
		}
	}

	var qs []*survey.Question
	if cli.Commits {
		rawCommits, err := git.Commits()
		if err != nil {
			panic(err)
		}

		qs = []*survey.Question{
			{
				Name: "target",
				Prompt: &survey.Select{
					Message: "Select a commit",
					Options: slice.Map[[]string, string](rawCommits, func(commit []string) string {
						commitId, message := slice.Shift[string](commit)

						return fmt.Sprintf("%s %s", highlight(commitId), strings.Join(message, " "))
					}),
				},
			},
		}
	} else {
		branches, err := git.Branch()

		if err != nil {
			panic(err)
		}

		qs = []*survey.Question{
			{
				Name: "target",
				Prompt: &survey.Select{
					Message: "Select a branch",
					Options: branches,
					Default: git.GetCurrentBranch(),
				},
			},
		}
	}

	answer := struct {
		Target string `survey:"target"`
	}{}

	err := survey.Ask(qs, &answer)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	git.Checkout(answer.Target, false)
}
