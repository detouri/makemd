package cli

import "github.com/spf13/cobra"

var (
	version = "0.4.2"
	rootCmd = &cobra.Command{
		Use:           "makemd",
		Short:         "Generate and maintain purposeful README.md files",
		Long:          "makemd creates fit-for-purpose README templates for apps, libraries, packages, CLIs, and internal tools.",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Version = version
}
