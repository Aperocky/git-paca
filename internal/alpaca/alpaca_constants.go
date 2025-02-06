package alpaca

const (
	SummaryCommmand  = "summarize"
	SummaryPrompt    = "You are a Git expert. Please provide a clear, focused summary of the important changes shown in the GIT OUTPUT, highlighting any crucial modifications and their potential impact, no longer than 100 words, only output the summary: "
	ReviewCommand    = "review"
	ReviewPrompt     = "You are a senior software engineer. Review the GIT OUTPUT and highlight both critical issues and notable improvements needed, focusing on security, performance, and maintainability: "
	CommitMsgCommand = "commitmsg"
	CommitMsgPrompt  = "You are a Git expert. Given the provided GIT OUTPUT, write a clear and descriptive commit message following conventional commits format, focusing on the what and why of the changes, no longer than 50 words, only output the commit message: "
)

var promptMap = map[string]string{
	SummaryCommmand:  SummaryPrompt,
	ReviewCommand:    ReviewPrompt,
	CommitMsgCommand: CommitMsgPrompt,
}
