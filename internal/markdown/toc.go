package markdown

import (
	"fmt"
	"strings"
)

func TOCFromSections(sections []Section) string {
	var items []string
	for _, sec := range sections {
		if sec.Hidden {
			continue
		}
		if strings.TrimSpace(sec.Title) == "" {
			continue
		}
		indent := strings.Repeat("	", max(sec.Level-2, 0))
		items = append(items, fmt.Sprintf("%s- [%s](#%s)", indent, sec.Title, Anchor(sec.Title)))
	}
	return strings.Join(items, "\n")
}
