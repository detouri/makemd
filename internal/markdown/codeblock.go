package markdown

import (
	"fmt"
)

func CodeBlock(lang, code string) string {
	return fmt.Sprintf("```%s\n%s\n```\n", lang, code)
}
