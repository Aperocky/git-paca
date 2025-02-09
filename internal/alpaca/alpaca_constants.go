package alpaca

const (
	GitSystemPrompt = `You are an git analysis assistant, understanding that (+) indicates additions, (-) indicates deletions, and surrounding lines provide context for the changes.`
	GitPrimer       = `Git changes are shown in a diff format:
- Lines starting with + show code that was added
- Lines starting with - show code that was removed
- Lines without +/- provide context around the changes
- Files being modified are marked with 'diff --git'
- File metadata includes @@ markers showing line numbers
`

	SummaryCommand = "summarize"
	SummaryPrompt  = `Here are your instruction:
Output only a summary: As a Git expert, analyze the following GIT OUTPUT and provide a focused summary (max 150 words) highlighting key modifications and their potential impact.`

	ReviewCommand = "review"
	ReviewPrompt  = `Here are your instruction:
With a focus on code quality, analyze the following GIT OUTPUT critically. Focus on identifying genuine concerns across these categories:

- Critical Issues (bugs and breaking changes)
- Performance Bottlenecks
- Maintainability Risks

Only report categories where you identify specific, actionable issues. For each identified issue:
- Describe the problem
- Explain its potential impact
- Suggest a concrete improvement

You can potentially find multiple issues or none in each category, if none is found, omit that category entirely.`

	CommitMsgCommand = "commitmsg"
	CommitMsgPrompt  = `Here are your instruction:
Generate a conventional commit message (<type>(?<scope>): <description>) using feat/fix/docs/style/refactor/test/chore/perf, imperative mood, make concise but complete description of the changes (max 50 words, output only commit message).`
)

var promptMap = map[string]string{
	SummaryCommand:   SummaryPrompt,
	ReviewCommand:    ReviewPrompt,
	CommitMsgCommand: CommitMsgPrompt,
}
