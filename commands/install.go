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

		if err := utils.RunCommand("sudo", "apt", "update"); err != nil {
			fmt.Printf("❌ Failed to update package lists: %v\n", err)
			return
		}

		if err := utils.RunCommand("sudo", "apt", "install", "-y", "curl", "wget", "unzip", "docker.io", "docker-compose"); err != nil {
			fmt.Printf("❌ Failed to install dependencies: %v\n", err)
			return
		}

		fmt.Println("📦 Cloning Minecraft Server repo...")
		if err := utils.RunCommand("git", "clone", serverRepoUrl, "~/minecraft"); err != nil {
			fmt.Printf("❌ Failed to clone server repository: %v\n", err)
			return
		}

		fmt.Println("🚀 Running initial setup...")
		if err := utils.RunCommand("mc-rubrion-cli", "start"); err != nil {
			fmt.Printf("❌ Failed to start the server after installation: %v\n", err)
			return
		}

		fmt.Println("✅ Installation complete! Use 'mc-rubrion-cli start' to launch the server.")
	},
}
