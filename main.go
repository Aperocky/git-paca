package main

import (
	"flag"
	"log"

	"github.com/Aperocky/git-paca/internal/alpaca"
	"github.com/Aperocky/git-paca/internal/config"
	"github.com/Aperocky/git-paca/internal/gitcmd"
	"github.com/Aperocky/git-paca/internal/parser"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error reading configuration: %v", err)
	}

	flag.Parse()
	args := flag.Args()

	cmdArgs, err := parser.ParseArguments(args)
	if err != nil {
		log.Fatalf("Cannot parse provided arguments: %v", err)
	}

	result, err := gitcmd.RunGitCmd(cmdArgs.GitCommands)
	if err != nil {
		log.Fatalf("Could not run git command: %v", err)
	}

	err = alpaca.AlpacaStream(config, string(result), cmdArgs)
	if err != nil {
		log.Fatalf("Error executing git-paca against ollama: %v", err)
	}
}
