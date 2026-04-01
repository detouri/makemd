package markdown

import "strings"

type Document struct {
	Title    string
	Preamble []string
	Sections []Section
}

func (d Document) Render() string {
	var parts []string
	if strings.TrimSpace(d.Title) != "" {
		parts = append(parts, "# "+strings.TrimSpace(d.Title))
	}
	if len(d.Preamble) > 0 {
		preamble := strings.TrimSpace(strings.Join(d.Preamble, "\n"))
		if preamble != "" {
			parts = append(parts, preamble)
		}
	}

	for _, sec := range d.Sections {
		if sec.Hidden {
			continue
		}
		body := strings.TrimSpace(sec.Body)
		if sec.Comment != "" {
			if body != "" {
				body += "\n"
			}
			body += toComment(sec.Comment)
		}
		title := strings.TrimSpace(sec.Title)
		if title == "" {
			if body != "" {
				parts = append(parts, body)
			}
			continue
		}
		heading := strings.Repeat("#", normaliseLevel(sec.Level)) + " " + title
		if body == "" {
			parts = append(parts, heading)
			continue
		}
		parts = append(parts, heading+"\n\n"+body)
	}
	return strings.TrimSpace(strings.Join(parts, "\n\n")) + "\n"
}
