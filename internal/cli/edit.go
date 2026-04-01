package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/detouri/makemd/internal/markdown"
	"github.com/spf13/cobra"
)

type multiFlag []string

// Type implements pflag.Value.
func (m *multiFlag) Type() string {
	return ""
}

func (m *multiFlag) String() string {
	return strings.Join(*m, ",")
}

func (m *multiFlag) Set(v string) error {
	*m = append(*m, v)
	return nil
}

var editFile string
var editSet multiFlag
var editRemove multiFlag
var editAppend multiFlag
var editTitle string
var editTOC bool

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringVar(&editFile, "file", "README.md", "README file to edit")
	editCmd.Flags().Var(&editSet, "set", "set section content by title: title=<markdown>")
	editCmd.Flags().Var(&editRemove, "remove", "remove section by title")
	editCmd.Flags().Var(&editAppend, "append", "append section by title: title=<markdown>")
	editCmd.Flags().StringVar(&editTitle, "title", "", "replace the H1 title")
	editCmd.Flags().BoolVar(&editTOC, "toc", false, "rebuild table of contents")
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing README.md",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := os.ReadFile(editFile)
		if err != nil {
			return err
		}
		editor := markdown.NewEditor(string(data))
		if strings.TrimSpace(editTitle) != "" {
			editor.SetTitle(editTitle)
		}
		for _, op := range editSet {
			k, v, err := splitKV(op)
			if err != nil {
				return err
			}
			editor.SetSection(k, v)
		}
		for _, sec := range editRemove {
			editor.RemoveSection(sec)
		}
		for _, op := range editAppend {
			k, v, err := splitKV(op)
			if err != nil {
				return err
			}
			editor.AppendSection(k, v)
		}
		if editTOC {
			editor.UpsertTOC("Table of Contents")
		}
		if err := os.WriteFile(editFile, []byte(editor.Render()), 0o644); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "updated %s\n", editFile)
		return nil
	},
}

func splitKV(in string) (string, string, error) {
	parts := strings.SplitN(in, "=", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid format %q; expected title=<markdown>", in)
	}
	v := strings.ReplaceAll(parts[1], `\n`, "\n")
	v = strings.ReplaceAll(v, `\t`, "\t")
	return strings.TrimSpace(parts[0]), v, nil
}
