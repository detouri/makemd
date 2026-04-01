package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/detouri/makemd/internal/markdown"
	"github.com/spf13/cobra"
)

var badgeFile string
var badgeAdd multiFlag

func init() {
	rootCmd.AddCommand(badgesCmd)
	badgesCmd.Flags().StringVar(&badgeFile, "file", "README.md", "README file to update")
	badgesCmd.Flags().Var(&badgeAdd, "add", "badge in the form alt|imageURL|linkURL(optional)")
}

var badgesCmd = &cobra.Command{
	Use:   "badges",
	Short: "Insert badge markdown into an existing README.md",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(badgeAdd) == 0 {
			return fmt.Errorf("at least one --add badge is required")
		}
		data, err := os.ReadFile(badgeFile)
		if err != nil {
			return err
		}
		editor := markdown.NewEditor(string(data))

		badges := make([]string, 0, len(badgeAdd))
		for _, item := range badgeAdd {
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
		if err := os.WriteFile(badgeFile, []byte(editor.Render()), 0o644); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "updated badges in %s\n", badgeFile)
		return nil
	},
}
