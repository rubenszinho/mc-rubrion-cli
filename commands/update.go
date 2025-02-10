package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

type Release struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Check for and update to the latest version of mc-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔄 Checking for updates...")

		latestRelease, err := fetchLatestRelease()
		if err != nil {
			fmt.Println("❌ Failed to check for updates:", err)
			return
		}

		latestVersion := latestRelease.TagName
		if latestVersion == "" || latestVersion == CurrentVersion {
			fmt.Println("✅ mc-cli is up to date (v" + CurrentVersion + ")")
			return
		}

		fmt.Println("🆕 New version available:", latestVersion)
		fmt.Println("⬇️ Downloading update...")

		if err := updateBinary(latestRelease); err != nil {
			fmt.Println("❌ Failed to update mc-cli:", err)
			return
		}

		fmt.Println("✅ Update successful! Restart mc-cli to use the latest version.")
	},
}

func fetchLatestRelease() (*Release, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", repoOwner, repoName)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}

	return &release, nil
}

func updateBinary(release *Release) error {
	osName := runtime.GOOS
	var binaryName string

	switch osName {
	case "linux":
		binaryName = "mc-cli-linux"
	case "windows":
		binaryName = "mc-cli.exe"
	case "darwin":
		binaryName = "mc-cli-mac"
	default:
		return fmt.Errorf("unsupported OS: %s", osName)
	}

	var downloadURL string
	for _, asset := range release.Assets {
		if asset.BrowserDownloadURL != "" && asset.BrowserDownloadURL[len(asset.BrowserDownloadURL)-len(binaryName):] == binaryName {
			downloadURL = asset.BrowserDownloadURL
			break
		}
	}

	if downloadURL == "" {
		return fmt.Errorf("could not find a suitable binary for %s", osName)
	}

	resp, err := http.Get(downloadURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tempFile, err := os.CreateTemp("", "mc-cli-*")
	if err != nil {
		return err
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		return err
	}

	err = os.Rename(tempFile.Name(), "/usr/local/bin/mc-cli")
	if err != nil {
		return err
	}

	err = exec.Command("chmod", "+x", "/usr/local/bin/mc-cli").Run()
	if err != nil {
		return err
	}

	return nil
}
