package commands

import (
	"fmt"
	"mc-rubrion-cli/discord"
	"mc-rubrion-cli/utils"

	"github.com/spf13/cobra"
)

var AddModCmd = &cobra.Command{
	Use:   "add-mod <URL>",
	Short: "Install a mod from a URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		fmt.Println("📦 Downloading mod from:", url)
		utils.RunCommand("wget", "-P", "~/minecraft/mods", url)

		fmt.Println("✅ Mod installed! Restart the server using 'mc-rubrion-cli restart'.")

		discord.UpdateDiscordStatus()
	},
}
