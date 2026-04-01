package readme

import "github.com/detouri/makemd/internal/markdown"

func genetedBy() markdown.Section {
	sec := markdown.Section{
		Title:   "",
		Body:    "",
		Comment: "This markdown file was create by makemd",
	}

	return sec
}
