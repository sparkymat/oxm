package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "oxm",
	Short: "oxm is a source-based package manager",
	Long:  "oxm is a source-based package manager which installs supported packages into a user's home directory. It installs it's own version of gcc.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
