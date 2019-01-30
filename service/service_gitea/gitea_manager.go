package service_gitea

import (
	"github.com/go-gitea/go-sdk/gitea"
	"lemon-robot-server/sysinfo"
)

var clientObj *gitea.Client

func Client() *gitea.Client {
	if clientObj == nil {
		clientObj = gitea.NewClient(sysinfo.LrConfig().SecretHmacKeyword, sysinfo.LrConfig().SecretHmacKeyword)
	}
	return clientObj
}
