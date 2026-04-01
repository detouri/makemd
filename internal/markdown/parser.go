package markdown

import "strings"

func Parse(content string) Document {
	lines := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")
	var doc Document
	var preamble []string
	var curr *Section
	var bLines []string

	flush := func() {
		if curr == nil {
			return
		}
		curr.Body = strings.TrimSpace(strings.Join(bLines, "\n"))
		doc.Sections = append(doc.Sections, *curr)
		curr = nil
		bLines = nil
	}

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "# ") && doc.Title == "" && curr == nil && len(doc.Sections) == 0 {
			doc.Title = strings.TrimSpace(strings.TrimPrefix(trimmed, "# "))
			continue
		}
		if strings.HasPrefix(trimmed, "##") {
			flush()
			level := headingLevel(trimmed)
			title := strings.TrimSpace(trimmed[level+1:])
			curr = &Section{Level: level, Title: title}
			continue
		}
		if curr == nil {
			preamble = append(preamble, line)
			continue
		}
		bLines = append(bLines, line)
	}
	flush()

	preambleText := strings.TrimSpace(strings.Join(preamble, "\n"))
	if preambleText != "" {
		doc.Preamble = splitBlocks(preambleText)
	}
	return doc
}

func splitBlocks(s string) []string {
	parts := strings.Split(s, "\n\n")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func headingLevel(s string) int {
	level := 0
	for _, r := range s {
		if r == '#' {
			level++
			continue
		}
		break
	}
	if level < 2 {
		return 2
	}
	return level
}
