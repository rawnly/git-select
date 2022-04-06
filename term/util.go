package term

import (
	"os"
	"os/exec"
)

// RunCommand Execute a command ignoring output
func RunCommand(command string, args ...string) (string, error) {
	stdout, err := exec.Command(command, args...).CombinedOutput()

	if err != nil {
		return "", err
	}

	return string(stdout), nil
}

func RunOSCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
