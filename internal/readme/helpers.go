package readme

import (
	"fmt"
	"strings"
)

func applyDerivedDefaults(c *ProjectConfig) {
	c.Description = fallback(c.Description, "A clean, purpose built project description goes here.")
	c.Repo = fallback(c.Repo, slugify(c.Title))
	c.Module = fallback(c.Module, fmt.Sprintf("github.com/%s/%s", fallback(c.Owner, "your-org"), c.Repo))
	c.BinaryName = fallback(c.BinaryName, firstWord(slugify(c.Title)))
	c.InstallCommand = fallback(c.InstallCommand, "go build ./...")
	c.RunCommand = fallback(c.RunCommand, "go run .")
	c.TestCommand = fallback(c.TestCommand, "go test ./...")
	c.CoverageCommand = fallback(c.CoverageCommand, c.TestCommand+" -cover")
	c.DocsURL = fallback(c.DocsURL, "https://example.com/docs")
	c.DemoURL = fallback(c.DemoURL, "https://example.com/demo")
	c.IssuesURL = fallback(c.IssuesURL, fmt.Sprintf("https://github.com/%s/%s/issues", fallback(c.Owner, "your-org"), c.Repo))
	if c.Private {
		c.LicenseName = fallback(c.LicenseName, "Proprietary")
		c.Status = fallback(c.Status, "internal")
	} else {
		c.LicenseName = fallback(c.LicenseName, "MIT")
		c.Status = fallback(c.Status, "active")
	}
	c.Audience = fallback(c.Audience, "engineers evaluating or using the project")
	c.MinimumVersion = fallback(c.MinimumVersion, "1.24+")
	c.CIProvider = fallback(c.CIProvider, "github-actions")
}

func fallback(val, def string) string {
	if val != "" {
		return val
	}
	return def
}

func slugify(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	replacer := strings.NewReplacer(" ", "-", "_", "-", "/", "-", ":", "", ".", "-")
	s = replacer.Replace(s)
	for strings.Contains(s, "--") {
		s = strings.ReplaceAll(s, "--", "-")
	}
	return strings.Trim(s, "-")
}

func firstWord(s string) string {
	parts := strings.FieldsFunc(s, func(r rune) bool { return r == '-' || r == ' ' })
	if len(parts) == 0 {
		return "app"
	}
	return parts[0]
}

func urlEscape(s string) string {
	repl := strings.NewReplacer("+", "%2B", " ", "%20", "/", "%2F")
	return repl.Replace(s)
}
