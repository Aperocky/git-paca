package parser

import (
	"fmt"
	"os"
	"slices"

	"github.com/Aperocky/git-paca/internal/alpaca"
	"github.com/Aperocky/git-paca/internal/types"
)

const (
	HelpString = `git paca cli syntax:
git-paca <git arguments..> <summarize|review|commitmsg|custom> [extra_prompt]
Examples:
git-paca diff summarize
git-paca diff --cached commitmsg
git-paca show review
git-paca diff custom "Give me all the typos"`
)

func ParseArguments(args []string) (*types.CommandArguments, error) {
	if args[0] == "help" || args[0] == "--help" {
		fmt.Println(HelpString)
		os.Exit(0)
	}
	if len(args) < 2 {
		fmt.Println(HelpString)
		return nil, fmt.Errorf("not enough arguments")
	}

	pacaCommands := make([]string, 0, len(alpaca.PromptMap))
	for k := range alpaca.PromptMap {
		pacaCommands = append(pacaCommands, k)
	}

	// last argument is paca-command
	if slices.Contains(pacaCommands, args[len(args)-1]) {
		return &types.CommandArguments{
			GitCommands: args[:len(args)-1],
			PacaCommand: args[len(args)-1],
			ExtraPrompt: "",
		}, nil
	}

	if len(args) == 2 {
		fmt.Println(HelpString)
		return nil, fmt.Errorf("argument not recognized, must be one of %v", pacaCommands)
	}

	// second to last argument is paca-command
	if slices.Contains(pacaCommands, args[len(args)-2]) {
		return &types.CommandArguments{
			GitCommands: args[:len(args)-2],
			PacaCommand: args[len(args)-2],
			ExtraPrompt: args[len(args)-1],
		}, nil
	}

	fmt.Println(HelpString)
	return nil, fmt.Errorf("incorrect command arguments/syntax")
}
