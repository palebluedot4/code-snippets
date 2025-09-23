package reversal_test

import (
	"testing"

	. "github.com/palebluedot4/code-snippets/cookbook/sequence-manipulation/reversal"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "single character",
			input: "a",
			want:  "a",
		},
		{
			name:  "ASCII string with even length",
			input: "abcd",
			want:  "dcba",
		},
		{
			name:  "ASCII string with odd length",
			input: "abcde",
			want:  "edcba",
		},
		{
			name:  "Unicode string",
			input: "你好，世界",
			want:  "界世，好你",
		},
		{
			name:  "string with spaces",
			input: "hello world",
			want:  "dlrow olleh",
		},
		{
			name:  "string with emoji",
			input: "Gopher ✅",
			want:  "✅ rehpoG",
		},
		{
			name:  "palindrome string",
			input: "madam",
			want:  "madam",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseString(tt.input)
			if got != tt.want {
				t.Errorf("ReverseString(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
