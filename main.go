package main

import (
	"fmt"
	"os"

	"mc-rubrion-cli/commands"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mc-rubrion-cli",
		Short: "Manage Minecraft Server with Docker",
	}

	rootCmd.AddCommand(commands.InstallCmd)
	rootCmd.AddCommand(commands.StartCmd)
	rootCmd.AddCommand(commands.StopCmd)
	rootCmd.AddCommand(commands.RestartCmd)
	rootCmd.AddCommand(commands.BackupCmd)
	rootCmd.AddCommand(commands.UpdateCmd)
	rootCmd.AddCommand(commands.AddModCmd)
	rootCmd.AddCommand(commands.StatusCmd)
	rootCmd.AddCommand(commands.VersionCmd)
	rootCmd.AddCommand(commands.ScheduleBackupCmd)
	rootCmd.AddCommand(commands.UnscheduleBackupCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
