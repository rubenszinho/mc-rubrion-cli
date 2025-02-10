package commands

import (
	"fmt"
	"mc-rubrion-cli/discord"

	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check and update the Minecraft server status in Discord",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("📡 Checking server status...")
		if err := discord.UpdateDiscordStatus(); err != nil {
			fmt.Printf("Failed to update Discord status: %v\n", err)
		} else {
			fmt.Println("✅ Status message updated in Discord.")
		}
	},
}
