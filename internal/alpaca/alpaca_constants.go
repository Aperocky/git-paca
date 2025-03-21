package alpaca

const (
	SummaryCommand = "summarize"
	SummaryPrompt  = `Analyze the following git diff output and provide a focused summary (max 150 words) highlighting key modifications and their potential impact.`

	ReviewCommand = "review"
	ReviewPrompt  = `Analyze following git diff output and report genuine concerns in any of these areas:

1. Critical Issues: Bugs, breaking changes, security issues
2. Performance: Bottlenecks, inefficient patterns
3. Maintainability: Complex logic, poor abstractions
4. Simplicity: Unnecessary complexity, redundancy

Skip categories with no issues. Multiple issues per category may be reported.`

	CommitMsgCommand = "commitmsg"
	CommitMsgPrompt  = `Generate a conventional commit message (<type>(?<scope>): <description>) using feat/fix/docs/style/refactor/test/chore/perf, imperative mood, make concise but complete description of the following git diff output (max 50 words, output only commit message).`

	CustomCommand = "custom"
)

var PromptMap = map[string]string{
	SummaryCommand:   SummaryPrompt,
	ReviewCommand:    ReviewPrompt,
	CommitMsgCommand: CommitMsgPrompt,
	CustomCommand:    "",
}
