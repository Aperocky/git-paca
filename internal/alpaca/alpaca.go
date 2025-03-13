package alpaca

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Aperocky/git-paca/internal/types"
)

const (
	TokenMultiplier = 2.0
	CustomHeader    = "Analyze the following git diff with these instructions: "
)

type streamResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func createReq(config *types.PacaConfig, payload string, cmdArgs *types.CommandArguments) (*types.OllamaRequest, error) {
	var prompt string
	if cmdArgs.PacaCommand == CustomCommand {
		prompt = CustomHeader + cmdArgs.ExtraPrompt + "\n"
	} else {
		commandPrompt := PromptMap[cmdArgs.PacaCommand]
		prompt = commandPrompt + "\n" + payload
		if cmdArgs.ExtraPrompt != "" {
			prompt += "\n" + "with these additional instructions: " + cmdArgs.ExtraPrompt
		}
	}

	neededCtx := int(float64(CountTokens(prompt)) * TokenMultiplier)

	ollamaOptions := config.Options
	if neededCtx > ollamaOptions.NumCtx {
		if neededCtx > config.MaxCtx {
			ollamaOptions.NumCtx = config.MaxCtx
		} else {
			ollamaOptions.NumCtx = neededCtx
		}
	}

	if config.Verbose {
		fmt.Println(prompt)
		fmt.Printf("Prompt has %v tokens\n", CountTokens(prompt))
		fmt.Printf("Setting num_ctx to %v\n", ollamaOptions.NumCtx)
	}

	reqBody := &types.OllamaRequest{
		Model:   config.ModelName,
		Prompt:  prompt,
		Stream:  true,
		Options: ollamaOptions,
	}
	return reqBody, nil
}

func AlpacaStream(config *types.PacaConfig, payload string, cmdArgs *types.CommandArguments) error {
	reqBody, err := createReq(config, payload, cmdArgs)
	if err != nil {
		return err
	}

	jsonData, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", config.Url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("server returned error %d %s: %s", resp.StatusCode, resp.Status, string(body))
	}

	decoder := json.NewDecoder(resp.Body)

	for {
		var streamResp streamResponse
		if err := decoder.Decode(&streamResp); err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error decoding response: %w", err)
		}
		fmt.Print(streamResp.Response)

		if streamResp.Done {
			fmt.Println("\n\nGit Paca Complete")
			break
		}
	}
	return nil
}
