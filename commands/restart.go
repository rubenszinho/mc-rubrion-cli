package commands

import (
	"fmt"
	"mc-rubrion-cli/discord"
	"mc-rubrion-cli/utils"

	"github.com/spf13/cobra"
)

var RestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the Minecraft server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔄 Restarting Minecraft Server...")

		if err := utils.RunCommand("docker-compose", "down"); err != nil {
			fmt.Printf("❌ Failed to stop the server: %v\n", err)
			return
		}

		if err := utils.RunCommand("docker-compose", "up", "-d"); err != nil {
			fmt.Printf("❌ Failed to start the server after restart: %v\n", err)
			return
		}

		fmt.Println("✅ Minecraft Server Restarted!")

		if err := discord.UpdateDiscordStatus(); err != nil {
			fmt.Printf("Failed to update Discord status: %v\n", err)
		}

	},
}
