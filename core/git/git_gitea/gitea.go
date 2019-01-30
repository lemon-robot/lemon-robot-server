package git_gitea

import (
	"github.com/go-gitea/go-sdk/gitea"
	"lemon-robot-golang-commons/logger"
	"os"
)

type GitGitea struct {
	client *gitea.Client
}

func (git *GitGitea) Init(config map[string]string) {
	endpoint := config["endpoint"]
	git.client = gitea.NewClient(endpoint, config["token"])
	_, err := git.client.ListMyRepos()
	if err != nil {
		logger.Error("Error connecting git server["+endpoint+"]", err)
		os.Exit(1)
	}
	logger.Info("Successful connection to git server: " + endpoint)
}
