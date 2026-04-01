package readme

import (
	"strings"
	"testing"
)

func TestServiceGenerateRequiresTemplateAndTitle(t *testing.T) {
	svc := NewService()

	if _, err := svc.Generate(ProjectConfig{}); err == nil || !strings.Contains(err.Error(), "template is required") {
		t.Fatalf("expected missing template error, got %v", err)
	}

	if _, err := svc.Generate(ProjectConfig{Template: "cli"}); err == nil || !strings.Contains(err.Error(), "title is required") {
		t.Fatalf("expected missing title error, got %v", err)
	}
}

func TestServiceGenerateRejectsUnknownTemplate(t *testing.T) {
	svc := NewService()

	_, err := svc.Generate(ProjectConfig{
		Template: "unknown",
		Title:    "Demo",
	})
	if err == nil || !strings.Contains(err.Error(), `unknown template "unknown"`) {
		t.Fatalf("expected unknown template error, got %v", err)
	}
}

func TestServiceGenerateCLIAndAppTemplates(t *testing.T) {
	svc := NewService()

	cliDoc, err := svc.Generate(ProjectConfig{
		Template: "cli",
		Title:    "MakeMD",
		Owner:    "detouri",
		Repo:     "makemd",
	})
	if err != nil {
		t.Fatalf("generate cli template: %v", err)
	}
	for _, want := range []string{
		"# MakeMD",
		"## Table of Contents",
		"## Commands",
		"makemd --help",
	} {
		if !strings.Contains(cliDoc, want) {
			t.Fatalf("expected CLI doc to contain %q:\n%s", want, cliDoc)
		}
	}

	appDoc, err := svc.Generate(ProjectConfig{
		Template: "app",
		Title:    "Demo Service",
	})
	if err != nil {
		t.Fatalf("generate app template: %v", err)
	}
	for _, want := range []string{
		"# Demo Service",
		"## Configuration",
		"## Architecture",
		"## Troubleshooting",
	} {
		if !strings.Contains(appDoc, want) {
			t.Fatalf("expected app doc to contain %q:\n%s", want, appDoc)
		}
	}
}

func TestServiceGenerateLibraryAndPackageTemplates(t *testing.T) {
	svc := NewService()

	tests := []struct {
		name     string
		cfg      ProjectConfig
		contains []string
		excludes []string
	}{
		{
			name: "open-source-lib",
			cfg: ProjectConfig{
				Template: "open-source-lib",
				Title:    "Public Library",
			},
			contains: []string{"## Security", "## License", "## Roadmap"},
		},
		{
			name: "closed-source-lib",
			cfg: ProjectConfig{
				Template: "closed-source-lib",
				Title:    "Internal Library",
			},
			contains: []string{"## Compatibility", "## Support", "## Ownership"},
			excludes: []string{"## License"},
		},
		{
			name: "package",
			cfg: ProjectConfig{
				Template: "package",
				Title:    "Tiny Package",
			},
			contains: []string{"## Installation", "## API", "## Examples"},
		},
	}

	for _, tc := range tests {
		doc, err := svc.Generate(tc.cfg)
		if err != nil {
			t.Fatalf("%s: generate failed: %v", tc.name, err)
		}
		for _, want := range tc.contains {
			if !strings.Contains(doc, want) {
				t.Fatalf("%s: expected output to contain %q:\n%s", tc.name, want, doc)
			}
		}
		for _, unwanted := range tc.excludes {
			if strings.Contains(doc, unwanted) {
				t.Fatalf("%s: expected output not to contain %q:\n%s", tc.name, unwanted, doc)
			}
		}
	}
}
