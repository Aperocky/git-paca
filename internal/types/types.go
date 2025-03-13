package types

type PacaConfig struct {
	Url       string         `json:"url"`
	ModelName string         `json:"model"`
	MaxCtx    int            `json:"maxctx,omitempty"`
	Verbose   bool           `json:"verbose,omitempty"`
	Options   *OllamaOptions `json:"options"`
}

type OllamaRequest struct {
	Model   string         `json:"model"`
	Prompt  string         `json:"prompt"`
	System  string         `json:"system"`
	Stream  bool           `json:"stream"`
	Options *OllamaOptions `json:"options,omitempty"`
}

type OllamaOptions struct {
	NumCtx        int     `json:"num_ctx,omitempty"`
	Temperature   float64 `json:"temperature,omitempty"`
	TopP          float64 `json:"top_p,omitempty"`
	TopK          int     `json:"top_k,omitempty"`
	RepeatPenalty float64 `json:"repeat_penalty,omitempty"`
}

type CommandArguments struct {
	GitCommands []string
	PacaCommand string
	ExtraPrompt string
}
