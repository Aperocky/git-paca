package gitcmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RunGitCmd(command string) ([]byte, error) {
	// Find git root
	gitRoot, err := findGitRoot()
	if err != nil {
		return nil, err
	}

	args := strings.Fields(command)
	if len(args) == 0 {
		return nil, fmt.Errorf("empty git command")
	}

	cmd := exec.Command("git", args...)
	cmd.Dir = gitRoot
	return cmd.CombinedOutput()
}

// findGitRoot returns the path to the root directory of the git repository
func findGitRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %v", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("not in a git repository")
		}
		dir = parent
	}
}
