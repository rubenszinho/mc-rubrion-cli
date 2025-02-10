# mc-rubrion-cli
A **Go-based command-line tool** for installing, managing, and monitoring the Minecraft server at **mc.rubrion.com** (or any domain you prefer).  
This CLI automates tasks such as installing Docker, cloning the server repo, managing mods, performing backups (including optional scheduled backups), sending status updates to Discord, and self-updating from GitHub.

---

## ✨ Features
- **One-liner Installation**: Installs Docker + your server in `~/minecraft`.
- **Start/Stop/Restart**: Runs Docker Compose commands to manage the server container.
- **Backup / Schedule-Backup**: Copies your server world data & optionally creates a **cron job** for automatic backups.
- **Add Mods**: Downloads mod `.jar` files into the server’s `mods` folder.
- **Status Updates**: Sends real-time info (CPU, RAM, player count, etc.) to a single Discord message (edits in place).
- **Auto-Updates**: Checks for new CLI versions on GitHub and updates itself.
- **Version**: Display the current CLI version.

---

## 🚀 Quick Start

1. **Download the Precompiled Binary**

   Visit the [Releases](https://github.com/rubrion/mc-rubrion-cli/releases) page and download the binary for your OS:
   - `mc-rubrion-cli-linux`
   - `mc-rubrion-cli.exe` (Windows)
   - `mc-rubrion-cli-mac`

2. **Install the CLI**
   ```bash
   chmod +x mc-rubrion-cli-linux
   sudo mv mc-rubrion-cli-linux /usr/local/bin/mc-rubrion-cli
   ```
   _(Adjust filenames and paths as needed.)_

3. **Configure Environment Variables (Optional)**
   - **`DISCORD_WEBHOOK_URL`** → Your **Discord** webhook URL (only needed if you want to post status updates to Discord).  
   - **`SERVER_DNS`** → The domain you want to display in Discord (default: `mc.rubrion.com`).  
   - **`REPO_OWNER`, `REPO_NAME`, `MC_CLI_VERSION`** → If you’re building from source and want to self-update from GitHub, you can set these at build time or environment. In most precompiled releases, this is already baked in.

4. **Install & Setup the Server**
   ```bash
   mc-rubrion-cli install
   ```
   - Installs dependencies (`docker`, `docker-compose`).
   - Clones the [minecraft-server-docker](https://github.com/rubrion/minecraft-server-docker.git) repo into `~/minecraft`.
   - Starts the server automatically.

5. **Manage the Server**
   ```bash
   mc-rubrion-cli start
   mc-rubrion-cli stop
   mc-rubrion-cli restart
   mc-rubrion-cli status
   mc-rubrion-cli backup
   mc-rubrion-cli add-mod <URL-to-mod>
   mc-rubrion-cli update
   mc-rubrion-cli schedule-backup "0 * * * *"   # optional cron-based auto-backup
   mc-rubrion-cli version                       # shows the CLI version
   ```

---

## 🕹 CLI Commands

| Command                       | Description                                                                 |
|-------------------------------|-----------------------------------------------------------------------------|
| **install**                   | Install dependencies & clone the server repo                                |
| **start**                     | Start the Minecraft server (docker-compose up -d)                           |
| **stop**                      | Stop the Minecraft server (docker-compose down)                             |
| **restart**                   | Restart the server (down, then up)                                          |
| **logs** (optional)           | View `docker` logs for `mc_server` (live console)                           |
| **backup**                    | Backup the server world data (copies `mc_server:/data`)                     |
| **schedule-backup** `<cron>` | Adds a **cron job** on Linux to run `mc-rubrion-cli backup` at a set interval |
| **add-mod <URL>**            | Download a mod `.jar` into `~/minecraft/mods`                               |
| **update**                    | Check for a new CLI version & update the binary (self-updater)              |
| **status**                    | Gathers stats & updates a single Discord message (PATCH if existing)        |
| **version**                   | Shows the current CLI version                                               |

---

## 🔗 Discord Integration
1. **Set `DISCORD_WEBHOOK_URL`** in your environment (only if you want Discord updates):
   ```bash
   export DISCORD_WEBHOOK_URL="https://discord.com/api/webhooks/..."
   ```
2. **Run server commands** like `start`, `stop`, `backup`, etc. The CLI will:
   - **Create** or **update** a **single** message in Discord showing CPU usage, RAM usage, uptime, player count, etc.

---

## 🗄 Repository Structure

- **commands/**: Go files for each CLI command (install, start, stop, restart, backup, schedule-backup, etc.).  
- **discord/**: Functions for sending or updating the single Discord message.  
- **utils/**: Shared helpers (`execute.go` for running commands, `getenv.go` for reading env vars).  
- **main.go**: Wires everything together with [Cobra](https://github.com/spf13/cobra).

---

## 🏗 Building from Source

```bash
git clone https://github.com/rubrion/mc-rubrion-cli.git
cd mc-rubrion-cli
go build -o mc-rubrion-cli main.go
```
Then move `mc-rubrion-cli` into your `$PATH` (e.g., `/usr/local/bin`).

If you want **build-time injection** of version info (for self-update checks), do:
```bash
go build -o mc-rubrion-cli \
  -ldflags="\
    -X 'mc-rubrion-cli/commands.CurrentVersion=1.2.3'"
  main.go
```

---

## 📝 License
Open-sourced under the **GPL-3.0** license. See the [LICENSE](LICENSE) file for details.