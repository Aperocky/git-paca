package alpaca

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Aperocky/git-paca/internal/types"
)

type streamResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func createReq(config *types.PacaConfig, payload string, command string) (*types.OllamaRequest, error) {
	commandPrompt, exists := promptMap[command]
	if !exists {
		return nil, fmt.Errorf("the command %s does not exist in git-paca", command)
	}
	prompt := commandPrompt + "### GIT DIFF output ###" + payload

	neededCtx := int(float64(CountTokens(prompt)) * 1.5)
	ollamaOptions := config.Options
	if neededCtx > ollamaOptions.NumCtx {
		if neededCtx > config.MaxCtx {
			ollamaOptions.NumCtx = config.MaxCtx
		} else {
			ollamaOptions.NumCtx = neededCtx
		}
	}

	if config.Verbose {
		fmt.Printf("Setting num_ctx to %v\n", ollamaOptions.NumCtx)
	}

	reqBody := &types.OllamaRequest{
		Model:   config.ModelName,
		Prompt:  prompt,
		System:  GitSystemPrompt,
		Stream:  true,
		Options: ollamaOptions,
	}
	return reqBody, nil
}

func AlpacaStream(config *types.PacaConfig, payload string, command string) error {
	reqBody, err := createReq(config, payload, command)
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
