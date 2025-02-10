package commands

import (
	"fmt"
	"time"

	"mc-rubrion-cli/discord"
	"mc-rubrion-cli/utils"

	"github.com/spf13/cobra"
)

var BackupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup the Minecraft world data",
	Run: func(cmd *cobra.Command, args []string) {
		timestamp := time.Now().Format("20060102_150405")
		utils.RunCommand("mkdir", "-p", "~/minecraft/backups")

		backupPath := fmt.Sprintf("~/minecraft/backups/world_backup_%s", timestamp)
		fmt.Println("💾 Backing up Minecraft world...")
		utils.RunCommand("docker", "cp", "mc_server:/data", backupPath)
		fmt.Println("✅ Backup saved at", backupPath)

		discord.UpdateDiscordStatus()
	},
}
