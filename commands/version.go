package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		if CurrentVersion == "" {
			fmt.Println("Version: (dev build)")
		} else {
			fmt.Printf("%s\n", CurrentVersion)
		}
	},
}
