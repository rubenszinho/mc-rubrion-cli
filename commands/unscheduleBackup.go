package commands

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var UnscheduleBackupCmd = &cobra.Command{
	Use:   "unschedule-backup",
	Short: "Remove any mc-rubrion-cli backup lines from crontab",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS != "linux" {
			fmt.Println("Un-scheduling only supported on Linux.")
			return
		}

		out, err := exec.Command("bash", "-c", "crontab -l").Output()
		if err != nil {
			fmt.Println("❌ Unable to read crontab:", err)
			return
		}
		lines := strings.Split(string(out), "\n")
		newLines := []string{}
		for _, line := range lines {
			if !strings.Contains(line, "mc-rubrion-cli backup") {
				newLines = append(newLines, line)
			}
		}
		joined := strings.Join(newLines, "\n")

		cmdString := fmt.Sprintf("echo \"%s\" | crontab -", joined)
		err = exec.Command("bash", "-c", cmdString).Run()
		if err != nil {
			fmt.Printf("❌ Failed to remove backup lines from crontab: %v\n", err)
			return
		}
		fmt.Println("✅ Removed all mc-rubrion-cli backup lines from crontab.")
	},
}
