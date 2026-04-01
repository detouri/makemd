package readme

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/detouri/makemd/internal/markdown"
)

type CLI struct {
	verison string
	stdout  io.Writer
	stderr  io.Writer
	service Service
}

func NewCLI(version string, stdout, stderr io.Writer) CLI {
	return CLI{
		verison: version,
		stdout:  stdout,
		stderr:  stderr,
		service: NewService(),
	}
}

func (c CLI) Run(args []string) error {
	if len(args) == 0 {
		c.printRootUsage()
		return nil
	}

	cmd := args[0]
	params := args[1:]

	switch args[0] {
	case "new":
		return c.runNew(params)
	case "edit":
		return c.runEdit(params)
	case "badges":
		return c.runBadges(params)
	case "list":
		return c.runList()
	case "version", "--version", "-v":
		_, _ = fmt.Fprintln(c.stdout, c.verison)
		return nil
	case "help", "--help", "-h":
		c.printRootUsage()
		return nil
	default:
		return fmt.Errorf("unknown command %q", cmd)
	}
}

func (c CLI) printRootUsage() {
	_, _ = fmt.Fprintln(c.stdout, `makemd - generate and maintain purposeful README.md files

Usage:
  makemd <command> [options]

Commands:
  new       Create a new README.md from a fit-for-purpose template
  edit      Add, replace, remove, or append sections in an existing README.md
  badges    Insert badge markdown into an existing README.md
  list      List built-in templates
  version   Print version`)
}

func (c CLI) runList() error {
	templates := c.service.List()
	sort.SliceStable(templates, func(i, j int) bool { return templates[i].Name < templates[j].Name })
	_, _ = fmt.Fprintln(c.stdout, "Available templates:")
	for _, tpl := range templates {
		_, _ = fmt.Fprintf(c.stdout, "  - %-18s %s\n", tpl.Name, tpl.Summary)
	}
	return nil
}

func (c CLI) runNew(args []string) error {
	fs := flag.NewFlagSet("new", flag.ContinueOnError)
	fs.SetOutput(c.stderr)

	cfg := ProjectConfig{}
	var output string
	var force bool

	fs.StringVar(&cfg.Template, "template", "", "template name")
	fs.StringVar(&cfg.Title, "title", "", "project title")
	fs.StringVar(&cfg.Description, "description", "", "short project description")
	fs.StringVar(&cfg.Owner, "owner", "", "repository owner or organization")
	fs.StringVar(&cfg.Repo, "repo", "", "repository name")
	fs.StringVar(&cfg.Module, "module", "", "Go module path")
	fs.StringVar(&cfg.BinaryName, "binary", "", "binary or CLI name")
	fs.StringVar(&cfg.InstallCommand, "install", "", "install command")
	fs.StringVar(&cfg.RunCommand, "run", "", "run command")
	fs.StringVar(&cfg.TestCommand, "test", "", "test command")
	fs.StringVar(&cfg.CoverageCommand, "coverage", "", "coverage command")
	fs.StringVar(&cfg.DocsURL, "docs-url", "", "documentation URL")
	fs.StringVar(&cfg.DemoURL, "demo-url", "", "demo URL")
	fs.StringVar(&cfg.IssuesURL, "issues-url", "", "issues URL")
	fs.StringVar(&cfg.CIProvider, "ci", "github-actions", "ci provider")
	fs.StringVar(&cfg.LicenseName, "license", "MIT", "license name")
	fs.StringVar(&cfg.Audience, "audience", "", "target audience")
	fs.StringVar(&cfg.Status, "status", "", "project status")
	fs.StringVar(&cfg.MinimumVersion, "version", "", "minimum version")
	fs.BoolVar(&cfg.Private, "private", false, "mark as closed-source/internal")
	fs.StringVar(&output, "output", "README.md", "output file")
	fs.BoolVar(&force, "force", false, "overwrite existing file")

	if err := fs.Parse(args); err != nil {
		return err
	}

	content, err := c.service.Generate(cfg)
	if err != nil {
		return err
	}

	if !force {
		if _, err := os.Stat(output); err == nil {
			return fmt.Errorf("%s already exists; use --force to overwrite", output)
		}
	}
	if err := os.WriteFile(output, []byte(content), 0o644); err != nil {
		return err
	}
	_, _ = fmt.Fprintf(c.stdout, "wrote %s using template %s\n", output, cfg.Template)
	return nil
}

func (c CLI) runEdit(args []string) error {
	fs := flag.NewFlagSet("edit", flag.ContinueOnError)
	fs.SetOutput(c.stderr)
	var file string
	var setOps multiFlag
	var removeOps multiFlag
	var appendOps multiFlag
	var title string
	var toc bool

	fs.StringVar(&file, "file", "README.md", "README file to edit")
	fs.Var(&setOps, "set", "set section content by key: title=<markdown>")
	fs.Var(&removeOps, "remove", "remove section by title")
	fs.Var(&appendOps, "append", "append a new section: title=<markdown>")
	fs.StringVar(&title, "title", "", "replace the H1 title")
	fs.BoolVar(&toc, "toc", false, "rebuild table of contents")
	if err := fs.Parse(args); err != nil {
		return err
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	editor := markdown.NewEditor(string(data))
	if strings.TrimSpace(title) != "" {
		editor.SetTitle(title)
	}
	for _, op := range setOps {
		k, v, err := splitKV(op)
		if err != nil {
			return err
		}
		editor.SetSection(k, v)
	}
	for _, name := range removeOps {
		editor.RemoveSection(name)
	}
	for _, op := range appendOps {
		k, v, err := splitKV(op)
		if err != nil {
			return err
		}
		editor.AppendSection(k, v)
	}
	if toc {
		editor.UpsertTOC("Table of Contents")
	}

	return os.WriteFile(file, []byte(editor.Render()), 0o644)
}

func (c CLI) runBadges(args []string) error {
	fs := flag.NewFlagSet("badges", flag.ContinueOnError)
	fs.SetOutput(c.stderr)
	var file string
	var add multiFlag
	fs.StringVar(&file, "file", "README.md", "README file to update")
	fs.Var(&add, "add", "badge in the form alt|imageURL|linkURL(optional)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if len(add) == 0 {
		return errors.New("at least one --add badge is required")
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	editor := markdown.NewEditor(string(data))

	badges := make([]string, 0, len(add))
	for _, item := range add {
		parts := strings.Split(item, "|")
		switch len(parts) {
		case 2:
			badges = append(badges, markdown.Badge(parts[0], parts[1]))
		case 3:
			badges = append(badges, markdown.Badge(parts[0], parts[1], parts[2]))
		default:
			return fmt.Errorf("invalid badge format %q; expected alt|imageURL|linkURL(optional)", item)
		}
	}
	editor.InsertBadges(badges)
	return os.WriteFile(file, []byte(editor.Render()), 0o644)
}

type multiFlag []string

func (m *multiFlag) String() string { return strings.Join(*m, ",") }
func (m *multiFlag) Set(v string) error {
	*m = append(*m, v)
	return nil
}

func splitKV(in string) (string, string, error) {
	parts := strings.SplitN(in, "=", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid format %q; expected title=<markdown>", in)
	}
	return strings.TrimSpace(parts[0]), normalizeMarkdownEscapes(parts[1]), nil
}

func normalizeMarkdownEscapes(s string) string {
	s = strings.ReplaceAll(s, `\n`, "\n")
	s = strings.ReplaceAll(s, `\t`, "\t")
	return s
}
