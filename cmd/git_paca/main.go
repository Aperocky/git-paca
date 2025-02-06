package main

import (
	"flag"
	"log"

	"github.com/Aperocky/git-paca/internal/alpaca"
	"github.com/Aperocky/git-paca/internal/config"
	"github.com/Aperocky/git-paca/internal/gitcmd"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error reading configuration: %v", err)
	}

	flag.Parse()
	args := flag.Args()

	gitCommand := args[0]
	pacaCommand := args[1]

	result, err := gitcmd.RunGitCmd(gitCommand)
	if err != nil {
		log.Fatalf("Could not run git command: %v", err)
	}

	err = alpaca.AlpacaStream(config, string(result), pacaCommand)
	if err != nil {
		log.Fatalf("Paca not connect to llama: %v", err)
	}
}
