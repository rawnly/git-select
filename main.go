package main

import (
	"fmt"
	"os/exec"
)


func gitBranch() error {
	out, err := exec.Command("git", "branch").Output()

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	fmt.Printf("The data is: %s\n", out)

	return nil
}

func gitCheckout() {

}

func main() {
	err := gitBranch()

	if err != nil {
		panic(err)
	}
}
