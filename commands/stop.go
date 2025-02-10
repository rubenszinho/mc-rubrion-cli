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
		utils.RunCommand("docker-compose", "down")
		fmt.Println("✅ Minecraft Server Stopped!")

		discord.UpdateDiscordStatus()
	},
}
