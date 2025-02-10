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
		utils.RunCommand("docker-compose", "down")
		utils.RunCommand("docker-compose", "up", "-d")
		fmt.Println("✅ Minecraft Server Restarted!")

		discord.UpdateDiscordStatus()
	},
}
