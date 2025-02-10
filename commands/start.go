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
		utils.RunCommand("docker-compose", "up", "-d")
		fmt.Println("✅ Minecraft Server Started!")

		discord.UpdateDiscordStatus()
	},
}
