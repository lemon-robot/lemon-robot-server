package gitter

import (
	"lemon-robot-server/core/gitter/git_gitea"
	"lemon-robot-server/entity"
)

/**
git标准接口
*/
type Standard interface {
	Init(map[string]string)

	TaskCreate(task *entity.Task) error
	TaskDelete(task entity.Task) error
	RepoContain(task entity.Task) bool
	RepoUrl(task entity.Task) (string, error)
}

var typesObj map[string]Standard

func SupportedTypes() map[string]Standard {
	if typesObj == nil {
		typesObj = make(map[string]Standard)
		typesObj["gitea"] = new(git_gitea.GitGitea)
	}
	return typesObj
}
