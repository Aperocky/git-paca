package gitcmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
)

const (
	GitCommandDiff = "diff"
	GitCommandShow = "show"
)

// Currently only supporting diff and show
// For review/commitmsg and summary
var supportedGitCommands = []string{
	GitCommandDiff,
	GitCommandShow,
}

func RunGitCmd(args []string) ([]byte, error) {
	// Find git root
	gitRoot, err := findGitRoot()
	if err != nil {
		return nil, err
	}

	if !slices.Contains(supportedGitCommands, args[0]) {
		log.Fatalf("Only supporting git commands %v", supportedGitCommands)
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
