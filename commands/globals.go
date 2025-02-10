package commands

import "mc-rubrion-cli/utils"

var (
	repoName      = utils.GetEnv("REPO_NAME", "mc-rubrion-cli")
	repoOwner     = utils.GetEnv("REPO_OWNER", "rubrion")
	serverRepoUrl = utils.GetEnv("SERVER_REPO_URL", "https://github.com/rubenszinho/mc-rubrion-server.git")
)

var CurrentVersion string
