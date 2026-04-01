package readme

import (
	"strings"
	"testing"
)

func TestApplyDerivedDefaultsPublicProject(t *testing.T) {
	cfg := ProjectConfig{
		Title: "My Tool",
		Owner: "detouri",
	}

	applyDerivedDefaults(&cfg)

	if cfg.Description == "" {
		t.Fatal("expected description default")
	}
	if cfg.Repo != "my-tool" {
		t.Fatalf("expected repo slug, got %q", cfg.Repo)
	}
	if cfg.Module != "github.com/detouri/my-tool" {
		t.Fatalf("unexpected module: %q", cfg.Module)
	}
	if cfg.BinaryName != "my" {
		t.Fatalf("unexpected binary name: %q", cfg.BinaryName)
	}
	if cfg.LicenseName != "MIT" {
		t.Fatalf("unexpected public license: %q", cfg.LicenseName)
	}
	if cfg.Status != "active" {
		t.Fatalf("unexpected status: %q", cfg.Status)
	}
	if cfg.MinimumVersion != "1.24+" {
		t.Fatalf("unexpected minimum version: %q", cfg.MinimumVersion)
	}
}

func TestApplyDerivedDefaultsPrivateProject(t *testing.T) {
	cfg := ProjectConfig{
		Title:   "Internal Library",
		Private: true,
	}

	applyDerivedDefaults(&cfg)

	if cfg.LicenseName != "Proprietary" {
		t.Fatalf("unexpected private license: %q", cfg.LicenseName)
	}
	if cfg.Status != "internal" {
		t.Fatalf("unexpected private status: %q", cfg.Status)
	}
}

func TestDefaultBadges(t *testing.T) {
	cfg := ProjectConfig{
		Owner:       "detouri",
		Repo:        "makemd",
		CIProvider:  "github-actions",
		LicenseName: "MIT",
		IssuesURL:   "https://github.com/detouri/makemd/issues",
		Status:      "active",
	}

	badges := DefaultBadges(cfg)
	joined := strings.Join(badges, " ")

	for _, want := range []string{
		"https://img.shields.io/github/actions/workflow/status/detouri/makemd/ci.yml?branch=main",
		"https://github.com/detouri/makemd/actions/workflows/ci.yml?query=branch%3Amain",
		"https://img.shields.io/github/issues/detouri/makemd",
		"https://img.shields.io/badge/status-active-blue",
		"https://img.shields.io/badge/license-MIT-green",
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("expected badges to contain %q:\n%s", want, joined)
		}
	}

	if strings.Contains(joined, "sheild") {
		t.Fatalf("unexpected typo in badge urls:\n%s", joined)
	}
}
