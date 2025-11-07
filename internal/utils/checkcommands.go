package utils

import (
	"fmt"
	"os/exec"
)

func CheckGitInstalled() error {
	cmd := exec.Command("git", "--version")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git is not installed or not available in PATH")
	}
	return nil
}
