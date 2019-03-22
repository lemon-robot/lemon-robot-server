package git_gitea

import (
	"github.com/go-gitea/go-sdk/gitea"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-server/entity"
	"os"
)

type GitGitea struct {
	client         *gitea.Client
	gitCurrentUser *gitea.User
}

func (gitter *GitGitea) Init(config map[string]string) {
	endpoint := config["endpoint"]
	gitter.client = gitea.NewClient(endpoint, config["token"])
	_, err := gitter.client.ListMyRepos()
	if err != nil {
		logger.Error("Error connecting gitter server["+endpoint+"]", err)
		os.Exit(1)
	}
	currentUser, err := gitter.client.GetMyUserInfo()
	if err != nil {
		logger.Error("Read current gitter user info error", err)
	}
	gitter.gitCurrentUser = currentUser
	logger.Info("Successful connection to gitter server: " + endpoint + ", user as: " + gitter.gitCurrentUser.UserName)
}

func (gitter *GitGitea) TaskCreate(task *entity.Task) error {
	_, orgErr := gitter.client.GetOrg(task.BelongNamespace.NamespaceTag)
	if orgErr != nil {
		logger.Error("Get gitter org error", orgErr)
		_, createOrgErr := gitter.client.AdminCreateOrg(gitter.gitCurrentUser.UserName, gitea.CreateOrgOption{
			UserName: task.BelongNamespace.NamespaceTag,
		})
		if createOrgErr != nil {
			logger.Error("Create gitter org error", createOrgErr)
		}
	}
	_, repoErr := gitter.client.CreateOrgRepo(task.BelongNamespace.NamespaceTag, gitea.CreateRepoOption{
		Name:    task.TaskTag,
		Private: true,
	})
	if repoErr != nil {
		logger.Error("Create gitter repo error", repoErr)
	}
	return repoErr
}

func (gitter *GitGitea) TaskDelete(task entity.Task) error {
	panic("implement me")
}

func (gitter *GitGitea) RepoContain(task entity.Task) bool {
	panic("implement me")
}

func (gitter *GitGitea) RepoUrl(task entity.Task) (string, error) {
	panic("implement me")
}
