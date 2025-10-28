package initialize

import (
	"web-scaffold/internal/core/config"
	"web-scaffold/pkg/encrypt"

	systemservice "web-scaffold/internal/service/system"
)

// @Author: yv1ing
// @Email:  me@yvling.cn
// @Date:   2025/10/28 14:50
// @Desc:	初始化系统用户

func InitSystemUser() error {
	username := "yv1ing"
	password := encrypt.Sha256String("123456", config.Config.SecretKey)
	name := "喻灵"
	email := "me@yvling.cn"
	phone := "13333333333"
	avatar := "https://avatars.githubusercontent.com/u/191813682"

	return systemservice.CreateUser(username, password, name, email, phone, avatar)
}
