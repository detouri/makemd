package cli

import (
	"fmt"
	"os"

	"github.com/detouri/makemd/internal/readme"
	"github.com/spf13/cobra"
)

var newOpts readme.ProjectConfig
var newOut string
var newForce bool

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVar(&newOpts.Template, "template", "", "template name")
	newCmd.Flags().StringVar(&newOpts.Title, "title", "", "project title")
	newCmd.Flags().StringVar(&newOpts.Description, "description", "", "short project description")
	newCmd.Flags().StringVar(&newOpts.Owner, "owner", "", "repository owner or organization")
	newCmd.Flags().StringVar(&newOpts.Repo, "repo", "", "repository name")
	newCmd.Flags().StringVar(&newOpts.Module, "module", "", "Go module path")
	newCmd.Flags().StringVar(&newOpts.BinaryName, "binary", "", "binary or CLI name")
	newCmd.Flags().StringVar(&newOpts.InstallCommand, "install", "", "install command")
	newCmd.Flags().StringVar(&newOpts.RunCommand, "run", "", "run command")
	newCmd.Flags().StringVar(&newOpts.TestCommand, "test", "", "test command")
	newCmd.Flags().StringVar(&newOpts.CoverageCommand, "coverage", "", "coverage command")
	newCmd.Flags().StringVar(&newOpts.DocsURL, "docs-url", "", "documentation URL")
	newCmd.Flags().StringVar(&newOpts.DemoURL, "demo-url", "", "demo URL")
	newCmd.Flags().StringVar(&newOpts.IssuesURL, "issues-url", "", "issues URL")
	newCmd.Flags().StringVar(&newOpts.CIProvider, "ci", "github-actions", "ci provider")
	newCmd.Flags().StringVar(&newOpts.LicenseName, "license", "MIT", "license name")
	newCmd.Flags().StringVar(&newOpts.Audience, "audience", "", "target audience")
	newCmd.Flags().StringVar(&newOpts.Status, "status", "", "project status")
	newCmd.Flags().StringVar(&newOpts.MinimumVersion, "go", "", "minimum Go version")
	newCmd.Flags().BoolVar(&newOpts.Private, "private", false, "mark as closed-source/internal")
	newCmd.Flags().StringVarP(&newOut, "output", "o", "README.md", "output file")
	newCmd.Flags().BoolVar(&newForce, "force", false, "overwrite existing file")

	_ = newCmd.MarkFlagRequired("template")
	_ = newCmd.MarkFlagRequired("title")
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new README.md from a template",
	RunE: func(cmd *cobra.Command, args []string) error {
		svc := readme.NewService()
		content, err := svc.Generate(newOpts)
		if err != nil {
			return err
		}
		if !newForce {
			if _, err := os.Stat(newOut); err == nil {
				return fmt.Errorf("%s already exists; use --force to overwirte", newOut)
			}
		}

		if err := os.WriteFile(newOut, []byte(content), 0o644); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "wrote %s using template %s\n", newOut, newOpts.Template)
		return nil
	},
}
