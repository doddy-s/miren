package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

func IsGitClean(dir string) bool {
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = dir

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		// Could not run git command, treat as not clean
		fmt.Println("Error:", err)
		return false
	}

	if out.Len() == 0 {
		// No changes, clean working directory
		return true
	}

	// There are changes or untracked files
	return false
}
