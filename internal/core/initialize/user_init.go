package initialize

import (
	"web-scaffold/internal/core/config"
	"web-scaffold/pkg/encrypt"

	systemmodel "web-scaffold/internal/model/system"
	systemrepository "web-scaffold/internal/repository/system"
)

// @Author: yv1ing
// @Email:  me@yvling.cn
// @Date:   2025/10/28 14:50
// @Desc:	初始化系统用户

func InitSystemUser() error {
	user := &systemmodel.User{
		Username: "yv1ing",
		Password: encrypt.Sha256String("123456", config.Config.SecretKey),
		Name:     "喻灵",
		Email:    "me@yvling.cn",
		Phone:    "13333333333",
		Avatar:   "https://avatars.githubusercontent.com/u/191813682",
		IsActive: true,
	}

	return systemrepository.CreateUser(user)
}
