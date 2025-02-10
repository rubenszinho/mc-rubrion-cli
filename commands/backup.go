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

		if err := utils.RunCommand("mkdir", "-p", "~/minecraft/backups"); err != nil {
			fmt.Printf("Failed to create backups directory: %v\n", err)
		}

		backupPath := fmt.Sprintf("~/minecraft/backups/world_backup_%s", timestamp)
		fmt.Println("💾 Backing up Minecraft world...")
		if err := utils.RunCommand("docker", "cp", "mc_server:/data", backupPath); err != nil {
			fmt.Printf("Failed to copy world data: %v\n", err)
		}
		fmt.Println("✅ Backup saved at", backupPath)

		if err := discord.UpdateDiscordStatus(); err != nil {
			fmt.Printf("Failed to update Discord status after backup: %v\n", err)
		}
	},
}
