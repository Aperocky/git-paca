package alpaca

const (
	SummaryCommand = "summarize"
	SummaryPrompt  = "Output only a summary: As a Git expert, analyze the GIT OUTPUT and provide a focused summary (max 150 words) highlighting key modifications and their potential impact."

	ReviewCommand = "review"
	ReviewPrompt  = "As a senior software engineer, review the GIT OUTPUT. Prioritize and list: 1) Critical Issues 2) Security Concerns 3) Performance Impacts 4) Maintainability Suggestions. Limit each section to 3 key points."

	CommitMsgCommand = "commitmsg"
	CommitMsgPrompt  = "Output only the commit message: As a Git expert, write a conventional commit message for the GIT OUTPUT. Focus on what changed and why (max 50 words)."
)

var promptMap = map[string]string{
	SummaryCommand:   SummaryPrompt,
	ReviewCommand:    ReviewPrompt,
	CommitMsgCommand: CommitMsgPrompt,
}
