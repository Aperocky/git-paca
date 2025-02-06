package alpaca

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Aperocky/git-paca/internal/config"
)

type ollamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type streamResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func AlpacaStream(config config.PacaConfig, payload string, command string) error {
	commandPrompt, exists := promptMap[command]
	if !exists {
		return fmt.Errorf("the command does not exist in git-paca")
	}
	prompt := commandPrompt + "### FOLLOWING IS GIT OUTPUT ###" + payload
	reqBody := ollamaRequest{
		Model:  config.ModelName,
		Prompt: prompt,
		Stream: true,
	}

	jsonData, _ := json.Marshal(reqBody)
	url := fmt.Sprintf("%s/api/generate", config.Url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
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
			fmt.Println("\n\n Git Paca Complete")
			break
		}
	}
	return nil
}
