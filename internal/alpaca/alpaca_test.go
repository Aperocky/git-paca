package alpaca

import "testing"

func TestCountTokens(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "single word",
			input:    "hello",
			expected: 2,
		},
		{
			name:     "short words",
			input:    "the cat sat",
			expected: 3,
		},
		{
			name:     "long word",
			input:    "supercalifragilistic",
			expected: 5, // 19 chars -> 5 tokens
		},
		{
			name:     "special characters",
			input:    "hello@world!",
			expected: 4,
		},
		{
			name:     "newlines",
			input:    "hello\nworld\n",
			expected: 4,
		},
		{
			name:     "git commit message",
			input:    "feat(api): add new endpoint\n\nCloses #123",
			expected: 14,
		},
		{
			name:     "code snippet",
			input:    "func main() {\n\treturn nil\n}",
			expected: 12,
		},
		{
			name:     "multiple spaces",
			input:    "hello    world",
			expected: 7, // 'hello' (2) + 'world' (1)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountTokens(tt.input)
			if got != tt.expected {
				t.Errorf("CountTokens(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
