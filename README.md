## git-paca

`git-paca` connects to locally ran LLM [Ollama](https://ollama.com/) to help you review, generate commit message and summaries on git changes and commits.

Installation: `go install github.com/Aperocky/git-paca@latest`

Running: `git-paca $git_command (summarize|review|commitmsg)`

For instance: `git-paca diff review`

### configuration

after installation and initial run, the configuration are stored in `~/.config/paca/paca-config.json`

Example (Default) Configuration:

```
{
  "url": "http://localhost:11434",
  "model": "qwen2.5-coder:14b-instruct-q6_K",
  "maxctx": 32768,
  "verbose": true,
  "options": {
    "num_ctx": 4096,
    "temperature": 0.3,
    "top_p": 0.9,
    "repeat_penalty": 1.1
  }
}
```

### Dogfooding Showcase

Commit messages for this package have come from itself.

*git-paca creating a commit message:*

```
$ ./git-paca diff --cached commitmsg
chore: Update RunGitCmd to accept []string and improve argument handling in main.go

Git Paca Complete
```
*git-paca summarize its initial commit:*

```
git-paca/bin$ ./git-paca show summarize
This code is a command-line tool named `git-paca` which integrates with Git and an AI model for automating commit messages, summarizing changes, or reviewing them based on the git repository's history. Here's a breakdown of its components:

1. **main.go**: The entry point of the application. It defines flags to accept two arguments from the user - the Git command (`gitCmd`) and the type of AI operation (`pacaCommand`).

2. **config/**: This package handles configuration management:
   - `config.go`: Contains methods for loading, creating, and storing a JSON configuration file located at ~/.config/paca/paca-config.json. The default values are defined in this file.

3. **gitcmd/**: This package executes Git commands within the context of the current working directory's git repository root.
   - `cmdrunner.go`: Implements functionality to run arbitrary Git commands and locate the nearest parent directory containing a `.git` folder (i.e., the root of the repository).

4. **alpaca/**: While not shown in your provided code snippet, based on naming conventions, this package likely interfaces with the ALPACA API or service using HTTP requests. It's responsible for sending Git logs and receiving processed responses from the AI model.

5. **utils.go**: Contains utility functions used throughout the application, such as `checkError` which prints an error message to stderr if there is one before exiting the program.

In summary, when you run this tool with a specific Git command (like "git diff") and specify what kind of AI operation you want ("review", "summarize", etc.), it will execute that Git command inside your repository, send its output to an AI model via HTTP, and then display the result back to you. The configuration can be customized by editing the paca-config.json file if needed.

Git Paca Complete
```
