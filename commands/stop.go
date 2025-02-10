package commands

import (
	"fmt"
	"mc-rubrion-cli/discord"
	"mc-rubrion-cli/utils"

	"github.com/spf13/cobra"
)

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the Minecraft server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🛑 Stopping Minecraft Server...")
		if err := utils.RunCommand("docker-compose", "down"); err != nil {
			fmt.Printf("❌ Failed to stop the server: %v\n", err)
			return
		}
		fmt.Println("✅ Minecraft Server Stopped!")

		if err := discord.UpdateDiscordStatus(); err != nil {
			fmt.Printf("Failed to update Discord status: %v\n", err)
		}
	},
}
