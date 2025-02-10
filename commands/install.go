package commands

import (
	"fmt"
	"mc-rubrion-cli/utils"

	"github.com/spf13/cobra"
)

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install and setup the Minecraft server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔧 Installing dependencies...")
		utils.RunCommand("sudo", "apt", "update")
		utils.RunCommand("sudo", "apt", "install", "-y", "curl", "wget", "unzip", "docker.io", "docker-compose")

		fmt.Println("📦 Cloning Minecraft Server repo...")
		utils.RunCommand("git", "clone", serverRepoUrl, "~/minecraft")

		fmt.Println("🚀 Running initial setup...")
		utils.RunCommand("mc-rubrion-cli", "start")

		fmt.Println("✅ Installation complete! Use 'mc-rubrion-cli start' to launch the server.")
	},
}
