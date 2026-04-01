package markdown

import "strings"

func Table(headers []string, rows [][]string) string {
	if len(headers) == 0 {
		return ""
	}
	cleanHeaders := make([]string, len(headers))
	separators := make([]string, len(headers))

	for i, h := range headers {
		cleanHeaders[i] = sanitizeCell(h)
		separators[i] = "---"
	}

	var lines []string
	lines = append(lines, "| "+strings.Join(cleanHeaders, " | ")+" |")
	lines = append(lines, "| "+strings.Join(separators, " | ")+" |")
	for _, row := range rows {
		cells := make([]string, len(headers))
		for i := range headers {
			if i < len(row) {
				cells[i] = sanitizeCell(row[i])
			} else {
				cells[i] = ""
			}
		}
		lines = append(lines, "| "+strings.Join(cells, " | ")+" |")
	}
	return strings.Join(lines, "\n")
}

func sanitizeCell(s string) string {
	return strings.ReplaceAll(strings.TrimSpace(s), "|", "\\|")
}
