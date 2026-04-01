package markdown

import (
	"fmt"
	"strings"
)

func Badge(alt, imageUrl string, linkUrl ...string) string {
	img := fmt.Sprintf("![%s](%s)", strings.TrimSpace(alt), strings.TrimSpace(imageUrl))
	if len(linkUrl) > 0 && strings.TrimSpace(linkUrl[0]) != "" {
		return fmt.Sprintf("[%s](%s)", img, strings.TrimSpace(linkUrl[0]))
	}
	return img
}
