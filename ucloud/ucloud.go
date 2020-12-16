package ucloud

import (
	"ucloudmanager/config"

	"github.com/ucloud/ucloud-sdk-go/services/uhost"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// BaseURL ucloud api接口
var BaseURL = "https://api.ucloud.cn"

// Uclient 初始化 ucloud 客户端
func Uclient(region string) *uhost.UHostClient {

	cfg := ucloud.NewConfig()
	cfg.Region = region
	cfg.BaseUrl = BaseURL

	c := config.Cfg.Cred
	cred := auth.NewCredential()
	cred.PublicKey = "" c.PublicKey
	cred.PrivateKey = c.PrivateKey

	uhostClient := uhost.NewClient(&cfg, &cred)
	return uhostClient
}
