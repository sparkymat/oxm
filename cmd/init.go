package cmd

import (
	"fmt"
	"os"

	"github.com/sparkymat/oxm/setup"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new oxm installation",
	Long:  "Sets up a new oxm installation in $HOME/.oxm.",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if err = setup.Init(); err == nil {
			fmt.Println("Successfully initialized. Try 'oxm status'.")
		} else {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}
