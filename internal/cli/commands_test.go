package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/detouri/makemd/internal/readme"
)

func restoreCLIState(t *testing.T) {
	t.Helper()

	prevNewOpts := newOpts
	prevNewOut := newOut
	prevNewForce := newForce
	prevEditFile := editFile
	prevEditSet := append(multiFlag(nil), editSet...)
	prevEditRemove := append(multiFlag(nil), editRemove...)
	prevEditAppend := append(multiFlag(nil), editAppend...)
	prevEditTitle := editTitle
	prevEditTOC := editTOC
	prevBadgeFile := badgeFile
	prevBadgeAdd := append(multiFlag(nil), badgeAdd...)

	t.Cleanup(func() {
		newOpts = prevNewOpts
		newOut = prevNewOut
		newForce = prevNewForce
		editFile = prevEditFile
		editSet = prevEditSet
		editRemove = prevEditRemove
		editAppend = prevEditAppend
		editTitle = prevEditTitle
		editTOC = prevEditTOC
		badgeFile = prevBadgeFile
		badgeAdd = prevBadgeAdd
		newCmd.SetOut(nil)
		editCmd.SetOut(nil)
		listCmd.SetOut(nil)
		badgesCmd.SetOut(nil)
	})
}

func TestSplitKV(t *testing.T) {
	key, value, err := splitKV(`Usage=Line one\nLine two\tTabbed`)
	if err != nil {
		t.Fatalf("splitKV returned error: %v", err)
	}
	if key != "Usage" {
		t.Fatalf("unexpected key: %q", key)
	}
	if value != "Line one\nLine two\tTabbed" {
		t.Fatalf("unexpected value: %q", value)
	}
}

func TestNewCommandWritesREADME(t *testing.T) {
	restoreCLIState(t)

	var out bytes.Buffer
	newCmd.SetOut(&out)

	newOpts = readme.ProjectConfig{
		Template:   "cli",
		Title:      "Demo CLI",
		Owner:      "detouri",
		Repo:       "makemd",
		BinaryName: "makemd",
	}
	newOut = filepath.Join(t.TempDir(), "README.md")
	newForce = true

	if err := newCmd.RunE(newCmd, nil); err != nil {
		t.Fatalf("new command failed: %v", err)
	}

	data, err := os.ReadFile(newOut)
	if err != nil {
		t.Fatalf("read generated README: %v", err)
	}
	if !strings.Contains(string(data), "# Demo CLI") {
		t.Fatalf("expected generated README, got:\n%s", string(data))
	}
	if !strings.Contains(out.String(), "wrote") {
		t.Fatalf("expected success output, got %q", out.String())
	}
}

func TestEditCommandUpdatesREADME(t *testing.T) {
	restoreCLIState(t)

	path := filepath.Join(t.TempDir(), "README.md")
	content := "# Demo\n\n## Usage\n\nOld body\n\n## Install\n\nOld install\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("seed README: %v", err)
	}

	editFile = path
	editSet = multiFlag{`Usage=New body`}
	editRemove = multiFlag{"Install"}
	editAppend = multiFlag{"Examples=Copy this example"}
	editTitle = "Updated Demo"
	editTOC = true

	if err := editCmd.RunE(editCmd, nil); err != nil {
		t.Fatalf("edit command failed: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read edited README: %v", err)
	}
	got := string(data)
	for _, want := range []string{
		"# Updated Demo",
		"## Table of Contents",
		"## Usage\n\nNew body",
		"## Examples\n\nCopy this example",
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("expected README to contain %q:\n%s", want, got)
		}
	}
	if strings.Contains(got, "## Install") {
		t.Fatalf("expected Install section removed:\n%s", got)
	}
}

func TestBadgesCommandRequiresAddAndWritesBadge(t *testing.T) {
	restoreCLIState(t)

	path := filepath.Join(t.TempDir(), "README.md")
	if err := os.WriteFile(path, []byte("# Demo\n"), 0o644); err != nil {
		t.Fatalf("seed README: %v", err)
	}

	badgeFile = path
	if err := badgesCmd.RunE(badgesCmd, nil); err == nil {
		t.Fatal("expected badges command to require at least one --add value")
	}

	badgeAdd = multiFlag{"CI|https://example.com/ci.svg|https://example.com/build"}
	if err := badgesCmd.RunE(badgesCmd, nil); err != nil {
		t.Fatalf("badges command failed: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read README: %v", err)
	}
	if !strings.Contains(string(data), "[![CI](https://example.com/ci.svg)](https://example.com/build)") {
		t.Fatalf("expected badge markdown in README:\n%s", string(data))
	}
}

func TestListCommandOutputsSortedTemplates(t *testing.T) {
	restoreCLIState(t)

	var out bytes.Buffer
	listCmd.SetOut(&out)

	if err := listCmd.RunE(listCmd, nil); err != nil {
		t.Fatalf("list command failed: %v", err)
	}

	output := out.String()
	expectedOrder := []string{
		"- app:",
		"- cli:",
		"- closed-source-lib:",
		"- open-source-lib:",
		"- package:",
	}
	lastIndex := -1
	for _, needle := range expectedOrder {
		idx := strings.Index(output, needle)
		if idx == -1 {
			t.Fatalf("expected %q in output:\n%s", needle, output)
		}
		if idx < lastIndex {
			t.Fatalf("expected sorted output, got:\n%s", output)
		}
		lastIndex = idx
	}
}
