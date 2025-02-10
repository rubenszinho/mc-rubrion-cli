package discord

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"mc-rubrion-cli/utils"
)

var (
	webhookURL    = utils.GetEnv("DISCORD_WEBHOOK_URL", "")
	messageIDFile = os.ExpandEnv("$HOME/minecraft/scripts/discord_message_id.txt")
	serverDNS     = utils.GetEnv("SERVER_DNS", "mc.rubrion.com")
)

type discordMessage struct {
	Content string `json:"content"`
}

func UpdateDiscordStatus() error {
	if webhookURL == "" {
		fmt.Println("❌ DISCORD_WEBHOOK_URL not set. Skipping Discord update.")
		return nil
	}

	cpuUsage := getCPUUsage()
	ramUsage := getRAMUsage()
	uptime := getUptime()
	playerCount := getPlayerCount()
	modCount := getModCount()

	message := fmt.Sprintf(
		`🟢 **Minecraft Server Status** 🟢
🌍 Address: %s
👥 Players Online: %s
📦 Modpack: %s Mods Installed
🔧 CPU Usage: %s%%
💾 RAM Usage: %s
🆙 Uptime: %s`,
		serverDNS,
		playerCount,
		modCount,
		cpuUsage,
		ramUsage,
		uptime,
	)

	data, err := os.ReadFile(messageIDFile)
	messageID := strings.TrimSpace(string(data))
	if err != nil || len(messageID) == 0 {
		return createDiscordMessage(message)
	} else {
		return editDiscordMessage(messageID, message)
	}
}

func createDiscordMessage(msg string) error {
	body, _ := json.Marshal(discordMessage{Content: msg})
	url := fmt.Sprintf("%s?wait=true", webhookURL)

	resp, err := http.Post(url, "application/json", strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("failed to send new Discord message, status: %d", resp.StatusCode)
	}

	respData, _ := io.ReadAll(resp.Body)
	messageID, err := parseMessageID(respData)
	if err != nil {
		return err
	}

	os.MkdirAll(os.ExpandEnv("$HOME/minecraft/scripts"), 0755)
	return os.WriteFile(messageIDFile, []byte(messageID), 0644)
}

func editDiscordMessage(messageID, msg string) error {
	body, _ := json.Marshal(discordMessage{Content: msg})
	client := &http.Client{}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/messages/%s", webhookURL, messageID), strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("failed to update Discord message, status: %d", resp.StatusCode)
	}

	return nil
}

func parseMessageID(data []byte) (string, error) {
	str := string(data)
	start := strings.Index(str, `"id":"`)
	if start == -1 {
		return "", fmt.Errorf("message ID not found in response")
	}
	start += 6
	end := strings.Index(str[start:], `"`)
	if end == -1 {
		return "", fmt.Errorf("invalid message ID format")
	}
	end = start + end
	return str[start:end], nil
}

func getCPUUsage() string {
	output, err := utils.CaptureCommand("bash", "-c", `top -bn1 | grep "Cpu(s)" | awk '{print $2 + $4}'`)
	if err != nil {
		return "N/A"
	}
	return strings.TrimSpace(output)
}

func getRAMUsage() string {
	output, err := utils.CaptureCommand("bash", "-c", `free -m | awk 'NR==2{printf "%s/%sMB (%.2f%%)", $3, $2, $3*100/$2 }'`)
	if err != nil {
		return "N/A"
	}
	return strings.TrimSpace(output)
}

func getUptime() string {
	output, err := utils.CaptureCommand("bash", "-c", `uptime -p | sed 's/up //'`)
	if err != nil {
		return "N/A"
	}
	return strings.TrimSpace(output)
}

func getPlayerCount() string {
	output, err := utils.CaptureCommand("bash", "-c", `docker exec mc_server rcon-cli list | grep -oP '\d+'`)
	if err != nil {
		return "N/A"
	}
	return strings.TrimSpace(output)
}

func getModCount() string {
	output, err := utils.CaptureCommand("bash", "-c", `ls ~/minecraft/mods 2>/dev/null | wc -l`)
	if err != nil {
		return "0"
	}
	return strings.TrimSpace(output)
}
