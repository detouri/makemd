package readme

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCLIRunListVersionAndHelp(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cli := NewCLI("1.2.3", &stdout, &stderr)

	if err := cli.Run([]string{"list"}); err != nil {
		t.Fatalf("list failed: %v", err)
	}
	if !strings.Contains(stdout.String(), "Available templates:") {
		t.Fatalf("expected list output, got:\n%s", stdout.String())
	}

	stdout.Reset()
	if err := cli.Run([]string{"version"}); err != nil {
		t.Fatalf("version failed: %v", err)
	}
	if strings.TrimSpace(stdout.String()) != "1.2.3" {
		t.Fatalf("unexpected version output: %q", stdout.String())
	}

	stdout.Reset()
	if err := cli.Run([]string{"help"}); err != nil {
		t.Fatalf("help failed: %v", err)
	}
	if !strings.Contains(stdout.String(), "makemd <command> [options]") {
		t.Fatalf("expected help output, got:\n%s", stdout.String())
	}
}

func TestCLIRunNewEditAndBadges(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cli := NewCLI("1.2.3", &stdout, &stderr)

	dir := t.TempDir()
	readmePath := filepath.Join(dir, "README.md")

	if err := cli.Run([]string{
		"new",
		"-template", "cli",
		"-title", "Demo CLI",
		"-owner", "detouri",
		"-repo", "makemd",
		"-binary", "makemd",
		"-output", readmePath,
	}); err != nil {
		t.Fatalf("new failed: %v\nstderr:\n%s", err, stderr.String())
	}

	if _, err := os.Stat(readmePath); err != nil {
		t.Fatalf("expected README to be written: %v", err)
	}

	if err := cli.Run([]string{
		"edit",
		"-file", readmePath,
		"-set", "Usage=Run `makemd --help` first.",
		"-append", "Examples=makemd run --config ./config.yaml",
		"-toc",
	}); err != nil {
		t.Fatalf("edit failed: %v\nstderr:\n%s", err, stderr.String())
	}

	if err := cli.Run([]string{
		"badges",
		"-file", readmePath,
		"-add", "CI|https://example.com/ci.svg|https://example.com/build",
	}); err != nil {
		t.Fatalf("badges failed: %v\nstderr:\n%s", err, stderr.String())
	}

	data, err := os.ReadFile(readmePath)
	if err != nil {
		t.Fatalf("read README: %v", err)
	}
	content := string(data)

	for _, want := range []string{
		"# Demo CLI",
		"[![CI](https://example.com/ci.svg)](https://example.com/build)",
		"Run `makemd --help` first.",
		"## Examples\n\nmakemd run --config ./config.yaml",
		"## Table of Contents",
	} {
		if !strings.Contains(content, want) {
			t.Fatalf("expected README to contain %q:\n%s", want, content)
		}
	}
}

func TestCLIRunRejectsUnknownCommand(t *testing.T) {
	cli := NewCLI("1.2.3", &bytes.Buffer{}, &bytes.Buffer{})

	if err := cli.Run([]string{"wat"}); err == nil || !strings.Contains(err.Error(), `unknown command "wat"`) {
		t.Fatalf("expected unknown command error, got %v", err)
	}
}
