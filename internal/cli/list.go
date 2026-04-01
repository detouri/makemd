package cli

import (
	"fmt"
	"sort"

	"github.com/detouri/makemd/internal/readme"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List avaliable templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		svc := readme.NewService()
		templates := svc.List()
		sort.SliceStable(templates, func(i, j int) bool {
			return templates[i].Name < templates[j].Name
		})
		for _, tpl := range templates {
			fmt.Fprintf(cmd.OutOrStdout(), "- %s: %s\n", tpl.Name, tpl.Summary)
		}
		return nil
	},
}
