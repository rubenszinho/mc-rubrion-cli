package commands

import (
	"fmt"
	"mc-rubrion-cli/discord"
	"mc-rubrion-cli/utils"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Minecraft server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Starting Minecraft Server...")
		if err := utils.RunCommand("docker-compose", "up", "-d"); err != nil {
			fmt.Printf("❌ Failed to start the server: %v\n", err)
			return
		}
		fmt.Println("✅ Minecraft Server Started!")

		if err := discord.UpdateDiscordStatus(); err != nil {
			fmt.Printf("Failed to update Discord status: %v\n", err)
		}
	},
}
