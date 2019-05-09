package service_gitea

import (
	"github.com/go-gitea/go-sdk/gitea"
	"lemon-robot-server/sysinfo"
)

var clientObj *gitea.Client

func Client() *gitea.Client {
	if clientObj == nil {
		clientObj = gitea.NewClient(sysinfo.LrServerConfig().SecretHmacKeyword, sysinfo.LrServerConfig().SecretHmacKeyword)
	}
	return clientObj
}
