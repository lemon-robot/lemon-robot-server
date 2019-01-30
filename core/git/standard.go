package git

import "lemon-robot-server/core/git/git_gitea"

/**
git标准接口
*/
type Standard interface {
	Init(map[string]string)
}

var typesObj map[string]Standard

func SupportedTypes() map[string]Standard {
	if typesObj == nil {
		typesObj = make(map[string]Standard)
		typesObj["gitea"] = new(git_gitea.GitGitea)
	}
	return typesObj
}
