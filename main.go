package main

import (
	"log"
	"web-scaffold/internal/core"
	"web-scaffold/internal/core/constant"
	"web-scaffold/internal/core/initialize"
	"web-scaffold/pkg/logger"
)

// @Author: yv1ing
// @Email:  me@yvling.cn
// @Date:   2025/10/28 11:07
// @Desc:   程序主入口

func main() {
	var err error

	// 初始化系统全局配置
	err = initialize.InitGlobalConfig("config.toml")
	if err != nil {
		log.Fatalf("[%d] %s\n", constant.CORE_INIT_CONF_ERROR, err.Error())
	}

	// 初始化系统全局日志
	logger.InitLogger("app.log", "debug")
	defer logger.Close()

	// 启动Web应用
	core.Start()
}
