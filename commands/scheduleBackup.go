package commands

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var ScheduleBackupCmd = &cobra.Command{
	Use:   "schedule-backup [cron_expression]",
	Short: "Schedule automatic backups via cron (Linux only). Example: mc-rubrion-cli schedule-backup \"0 * * * *\"",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cronExpr := args[0]

		if runtime.GOOS != "linux" {
			fmt.Println("❌ Auto backup scheduling is currently supported only on Linux.")
			return
		}
		cronCommand := fmt.Sprintf("(crontab -l 2>/dev/null; echo \"%s mc-rubrion-cli backup\") | crontab -", cronExpr)
		fmt.Printf("Scheduling automatic backups with cron expression: %q\n", cronExpr)
		fmt.Println("Command:", cronCommand)

		err := exec.Command("bash", "-c", cronCommand).Run()
		if err != nil {
			fmt.Printf("❌ Failed to update crontab: %v\n", err)
			return
		}

		fmt.Printf("✅ Successfully scheduled backups using cron: %s mc-rubrion-cli backup\n", cronExpr)
		fmt.Println("You can check your current crontab with: crontab -l")
	},
}
