package markdown

import "strings"

func Anchor(title string) string {
	s := strings.ToLower(strings.TrimSpace(title))
	replacer := strings.NewReplacer(" ", "-", "/", "", ":", "", ".", "", ",", "", "{`", "", "`", "")
	s = replacer.Replace(s)
	for strings.Contains(s, "--") {
		s = strings.ReplaceAll(s, "--", "-")
	}
	return strings.Trim(s, "-")
}
