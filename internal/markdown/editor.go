package markdown

import "strings"

type Editor struct {
	doc Document
}

func NewEditor(content string) Editor {
	return Editor{doc: Parse(content)}
}

func (e *Editor) SetTitle(title string) {
	e.doc.Title = strings.TrimSpace(title)
}

func (e *Editor) AddPreamble(block string) {
	block = strings.TrimSpace(block)
	if block == "" {
		return
	}
	e.doc.Preamble = append(e.doc.Preamble, block)
}

func (e *Editor) SetSection(title, body string) {
	normalised := strings.TrimSpace(strings.ToLower(title))
	for i, sec := range e.doc.Sections {
		if strings.ToLower(strings.TrimSpace(sec.Title)) == normalised {
			e.doc.Sections[i].Body = strings.TrimSpace(body)
			e.doc.Sections[i].Hidden = false
			return
		}
	}
	e.doc.Sections = append(e.doc.Sections, NewSection(title, body))
}

func (e *Editor) RemoveSection(title string) {
	normalised := strings.TrimSpace(strings.ToLower(title))
	filtered := e.doc.Sections[:0]
	for _, sec := range e.doc.Sections {
		if strings.ToLower(strings.TrimSpace(sec.Title)) == normalised {
			continue
		}
		filtered = append(filtered, sec)
	}
	e.doc.Sections = filtered
}

func (e *Editor) AppendSection(title, body string) {
	e.doc.Sections = append(e.doc.Sections, NewSection(title, body))
}

func (e *Editor) UpsertTOC(title string) {
	tocTitle := strings.TrimSpace(title)
	if tocTitle == "" {
		tocTitle = "Table of Contents"
	}
	filtered := make([]Section, 0, len(e.doc.Sections))
	for _, sec := range e.doc.Sections {
		if strings.EqualFold(strings.TrimSpace(sec.Title), tocTitle) {
			continue
		}
		filtered = append(filtered, sec)
	}
	toc := NewSection(tocTitle, TOCFromSections(filtered))
	e.doc.Sections = append([]Section{toc}, filtered...)
}

func (e *Editor) InsertBadges(badges []string) {
	line := strings.TrimSpace(strings.Join(badges, " "))
	if line == "" {
		return
	}
	if len(e.doc.Preamble) == 0 {
		e.doc.Preamble = append(e.doc.Preamble, line)
		return
	}
	e.doc.Preamble = append([]string{line}, e.doc.Preamble...)
}

func (e *Editor) Render() string {
	return e.doc.Render()
}
