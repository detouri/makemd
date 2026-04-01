package markdown

import (
	"fmt"
	"strings"
)

func Bullets(items ...string) string {
	out := make([]string, 0, len(items))
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		out = append(out, "- "+item)
	}
	return strings.Join(out, "\n")
}

func Numbered(items ...string) string {
	out := make([]string, 0, len(items))
	for n, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		out = append(out, fmt.Sprintf("%d. %s", n+1, item))
	}
	return strings.Join(out, "\n")
}
