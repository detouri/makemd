package markdown

import "strings"

func BlockQuote(text string) string {
	lines := strings.Split(strings.TrimSpace(text), "\n")
	for i, line := range lines {
		lines[i] = "> " + strings.TrimSpace(line)
	}
	return strings.Join(lines, "\n")
}
