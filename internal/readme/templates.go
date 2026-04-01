package readme

import (
	"fmt"

	"github.com/detouri/makemd/internal/markdown"
)

func Templates() map[string]Template {
	return map[string]Template{
		"open-source-lib": {
			Name:        "open-source-lib",
			Summary:     "Open source library with install, usage, API, testing, contributions and license sections.",
			Description: "Best for reusable public libraries and modules.",
			Build:       buildOpenSourceLib,
		},

		"closed-source-lib": {
			Name:        "closed-source-lib",
			Summary:     "Internal library focused on ownership, support, compatibility and usage within a project or organisation.",
			Description: "Best for internal packages and shared platform code.",
			Build:       buildClosedSourceLib,
		},

		"package": {
			Name:        "package",
			Summary:     "small package template with concise install, import, API, and example sectionns.",
			Description: "Best for small importable packages.",
			Build:       buildPackage,
		},

		"app": {
			Name:        "app",
			Summary:     "Application template with setup, config, running, architecture, deployement and troubleshooting sections.",
			Description: "Best for srvices, web apps, daemons and innternal products.",
			Build:       buildApp,
		},

		"cli": {
			Name:        "cli",
			Summary:     "CLI-orientated template with install methods, commands, examples, config and shell completetion.",
			Description: "Best for command-line tools.",
			Build:       buildCLI,
		},
	}
}

func buildOpenSourceLib(cfg ProjectConfig) (string, error) {
	cfg.IncludeBadges = true
	cfg.IncludeTOC = true
	cfg.IncludeFAQ = true
	cfg.IncludeRoadmap = true
	cfg.IncludeSecurity = true
	cfg.IncludeConfig = true
	cfg.IncludeAPI = true
	return renderLibDoc(cfg, false), nil
}

func buildClosedSourceLib(cfg ProjectConfig) (string, error) {
	cfg.Private = true
	cfg.IncludeBadges = true
	cfg.IncludeTOC = true
	cfg.IncludeAPI = true
	cfg.IncludeConfig = true
	return renderLibDoc(cfg, true), nil
}

func buildPackage(cfg ProjectConfig) (string, error) {
	// cfg.IncludeBadges = true
	cfg.IncludeTOC = true
	cfg.IncludeAPI = true
	doc := markdown.Document{Title: cfg.Title}
	if cfg.IncludeBadges {
		doc.Preamble = append(doc.Preamble, stringsJoin(DefaultBadges(cfg), " "))
	}
	doc.Preamble = append(doc.Preamble, markdown.BlockQuote(cfg.Description))

	doc.Sections = []markdown.Section{
		genetedBy(),
		markdown.NewSection("Installation", markdown.CodeBlock("bash", cfg.Module)),
		markdown.NewSection("Usage", ""),
		markdown.NewSection("API", "Document the stable surface area of the package and link to deeper docs if needed."),
		markdown.NewSection("Examples", "Point to `/examples` or include one minimal and one realistic example."),
		markdown.NewSection("Testing", markdown.CodeBlock("bash", cfg.TestCommand)),
		markdown.NewSection("License", fmt.Sprintf("This project is licensed under **%s**.", cfg.LicenseName)),
	}

	if cfg.IncludeTOC {
		doc.Sections = append([]markdown.Section{markdown.NewSection("Table of Content", markdown.TOCFromSections(doc.Sections))}, doc.Sections...)
	}
	return doc.Render(), nil
}

func buildApp(cfg ProjectConfig) (string, error) {
	cfg.IncludeBadges = true
	cfg.IncludeTOC = true
	cfg.IncludeRoadmap = true
	cfg.IncludeArch = true
	cfg.IncludeConfig = true

	doc := markdown.Document{Title: cfg.Title}
	if cfg.IncludeBadges {
		doc.Preamble = append(doc.Preamble, stringsJoin(DefaultBadges(cfg), " "))
	}
	doc.Preamble = append(doc.Preamble, markdown.BlockQuote(cfg.Description))

	sections := []markdown.Section{
		genetedBy(),
		markdown.NewSection("Features", markdown.Bullets(
			"Clear value proposition and project scope",
			"Fast path to local setup and first success",
			"Operational sections for maintainers and deployers",
		)),
		markdown.NewSection("Screenshots", "Add a screenshot, GIF, or demo link here to reduce time-to-understanding."),
		markdown.NewSection("Prerequisites", markdown.Bullets("Go "+cfg.MinimumVersion, "Git", "Required access or credentials")),
		markdown.NewSection("Installation", markdown.CodeBlock("bash", cfg.InstallCommand)),
		markdown.NewSection("Configuration", markdown.Table(
			[]string{"Key", "Required", "Default", "Description"},
			[][]string{{"PORT", "No", "8080", "Port the service listens on"}, {"LOG_LEVEL", "No", "info", "Logging verbosity"}, {"API_TOKEN", "Yes", "-", "Credential for upstream access"}},
		)),
		markdown.NewSection("Running", markdown.CodeBlock("bash", cfg.RunCommand)),
		markdown.NewSection("Usage", "Describe the most common workflow first and keep examples copy-pasteable."),
		markdown.NewSection("Architecture", markdown.Numbered(
			"Describe the major components.",
			"Explain request or data flow.",
			"List dependencies and storage.",
			"Link to deeper design docs.",
		)),
		markdown.NewSection("Deployment", markdown.Bullets(
			"Document environments and deployment steps",
			"List required secrets and rollback procedure",
			"Capture post-deploy verification checks",
		)),
		markdown.NewSection("Testing", markdown.CodeBlock("bash", cfg.TestCommand)),
		markdown.NewSection("Roadmap", markdown.Bullets(
			"[ ] Add production architecture diagram",
			"[ ] Add operational runbook links",
			"[ ] Add migration or release notes section",
		)),
		markdown.NewSection("Troubleshooting", markdown.Bullets(
			"Build fails: verify toolchain and access",
			"Runtime errors: confirm configuration and secrets",
			"Unexpected behavior: enable debug logs and capture repro steps",
		)),
		markdown.NewSection("Contributing", "Describe branch, test, and pull request expectations."),
		markdown.NewSection("License", fmt.Sprintf("This project is licensed under **%s**.", cfg.LicenseName)),
	}

	if cfg.IncludeTOC {
		sections = append([]markdown.Section{markdown.NewSection("Table of Contents", markdown.TOCFromSections(sections))}, sections...)
	}

	doc.Sections = sections
	return doc.Render(), nil
}

func buildCLI(cfg ProjectConfig) (string, error) {
	cfg.IncludeBadges = true
	cfg.IncludeTOC = true
	cfg.IncludeFAQ = true
	cfg.IncludeConfig = true
	doc := markdown.Document{Title: cfg.Title}
	if cfg.IncludeBadges {
		doc.Preamble = append(doc.Preamble, stringsJoin(DefaultBadges(cfg), " "))
	}
	doc.Preamble = append(doc.Preamble, markdown.BlockQuote(cfg.Description))

	sections := []markdown.Section{
		markdown.NewSection("Features", markdown.Bullets(
			"Focused command-line workflow",
			"Scriptable output and automation-friendly behavior",
			"Clear install and command reference",
		)),
		markdown.NewSection("Installation", markdown.CodeBlock("bash", cfg.InstallCommand)),
		markdown.NewSection("Quick Start", markdown.CodeBlock("bash", cfg.BinaryName+" --help")),
		markdown.NewSection("Commands", markdown.Table(
			[]string{"Command", "Description"},
			[][]string{{cfg.BinaryName + " --help", "Show global help"}, {cfg.BinaryName + " init", "Initialize configuration"}, {cfg.BinaryName + " validate", "Validate inputs"}, {cfg.BinaryName + " run", "Execute the main workflow"}},
		)),
		markdown.NewSection("Configuration", "List flags, environment variables, config file locations, and precedence rules."),
		markdown.NewSection("Shell Completion", markdown.CodeBlock("bash", fmt.Sprintf("%s completion bash > /etc/bash_completion.d/%s", cfg.BinaryName, cfg.BinaryName))),
		markdown.NewSection("Examples", markdown.CodeBlock("bash", cfg.BinaryName+" run --config ./config.yaml")),
		markdown.NewSection("Testing", markdown.CodeBlock("bash", cfg.TestCommand)),
		markdown.NewSection("FAQ", markdown.Bullets(
			"Who is this for?",
			"Is it production-ready?",
			"Where are the richer examples?",
		)),
		markdown.NewSection("License", fmt.Sprintf("This project is licensed under **%s**.", cfg.LicenseName)),
	}

	if cfg.IncludeTOC {
		sections = append([]markdown.Section{markdown.NewSection("Table of Contents", markdown.TOCFromSections(sections))}, sections...)
	}
	doc.Sections = sections
	return doc.Render(), nil
}

func renderLibDoc(cfg ProjectConfig, internal bool) string {
	doc := markdown.Document{Title: cfg.Title}
	if cfg.IncludeBadges {
		doc.Preamble = append(doc.Preamble, stringsJoin(DefaultBadges(cfg), " "))
	}
	doc.Preamble = append(doc.Preamble, markdown.BlockQuote(cfg.Description))

	sections := []markdown.Section{
		genetedBy(),
		markdown.NewSection("Features", markdown.Bullets(
			"Clear scope and integration path",
			"Practical install and usage guidance",
			"Structured sections for maintainers and consumers",
		)),
		markdown.NewSection("Installation", markdown.CodeBlock("bash", "go get "+cfg.Module)),
		markdown.NewSection("Quick Start", markdown.CodeBlock("go", fmt.Sprintf("import \"%s\"", cfg.Module))),
		markdown.NewSection("Usage", "Lead with one copy-pasteable example, then show a realistic integration path."),
		markdown.NewSection("API", "Document core types, functions, defaults, errors, and compatibility promises."),
		markdown.NewSection("Configuration", markdown.Table(
			[]string{"Key", "Required", "Default", "Description"},
			[][]string{{"LOG_LEVEL", "No", "info", "Logging verbosity"}, {"API_TOKEN", "Yes", "-", "Credential for upstream API access"}},
		)),
		markdown.NewSection("Examples", "Point to `/examples` or richer integration docs."),
		markdown.NewSection("Testing", markdown.CodeBlock("bash", cfg.TestCommand)),
	}

	if internal {
		sections = append(sections,
			markdown.NewSection("Compatibility", markdown.Bullets(
				"Minimum Language/runtime support version: "+cfg.MinimumVersion,
				"Document supported platforms and versioning policy",
				"List dependent teams or services if useful",
			)),
			markdown.NewSection("Support", markdown.Bullets(
				"Team channel: #team-support",
				"Escalation path: team@example.com",
				"Expected response time: define internal SLA",
			)),
			markdown.NewSection("FAQ", markdown.Bullets("Who is this for?", "What is the support policy?", "Where is the deep documentation?")),
			markdown.NewSection("Ownership", markdown.Bullets(
				"Owning team: Platform Engineering",
				"Maintainers: add names or handles",
				"Change policy: clarify review expectations",
			)),
		)
	} else {
		sections = append(sections,
			markdown.NewSection("Roadmap", markdown.Bullets("[ ] Add more examples", "[ ] Add migration guidance", "[ ] Document performance expectations")),
			markdown.NewSection("Contributing", "Describe branch, test, and pull request expectations."),
			markdown.NewSection("Security", "Document the private vulnerability disclosure path or link to SECURITY.md."),
			markdown.NewSection("FAQ", markdown.Bullets("Who is this for?", "Is it production-ready?", "Where are the full docs?")),
			markdown.NewSection("License", fmt.Sprintf("This project is licensed under **%s**.", cfg.LicenseName)),
		)
	}

	if cfg.IncludeTOC {
		sections = append([]markdown.Section{markdown.NewSection("Table of Contents", markdown.TOCFromSections(sections))}, sections...)
	}
	doc.Sections = sections
	return doc.Render()
}

func stringsJoin(items []string, sep string) string {
	if len(items) == 0 {
		return ""
	}
	out := items[0]
	for i := 1; i < len(items); i++ {
		out += sep + items[i]
	}
	return out
}
