package cmd

import (
	"fmt"
	"os"

	"github.com/sparkymat/oxm/path"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Checks the status of the oxm installation",
	Long:  "Checks if oxm is installed, and if it is, whether it is setup correctly",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var info os.FileInfo

		if info, err = os.Stat(path.Home()); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		if !info.IsDir() || (info.Mode().Perm()&(1<<(uint(7))) == 0) {
			fmt.Println("Error: Cannot write to $HOME/.oxm folder")
			os.Exit(1)
		}

		fmt.Println("Everything looks good")
	},
}
