package core

import (
	"fmt"
	"web-scaffold/internal/core/config"
	"web-scaffold/internal/core/initialize"
	"web-scaffold/internal/repository"
	"web-scaffold/pkg/logger"
)

// @Author: yv1ing
// @Email:  me@yvling.cn
// @Date:   2025/10/28 11:08
// @Desc:	Web应用入口

func Start() {
	var err error

	// 初始化数据库连接
	db, err := initialize.InitDatabase()
	if err != nil {
		logger.Error("Startup failed, error in creating a database connection: ", err)
		return
	}

	// 初始化数据仓储层
	err = repository.InitRepository(db)
	if err != nil {
		logger.Error("Startup failed, error in creating the data warehouse layer: ", err)
		return
	}

	// 初始化系统用户
	err = initialize.InitSystemUser()
	if err != nil {
		logger.Error("Startup failed, error in creating the system user: ", err)
		return
	}

	// 启动Web服务引擎
	eng := initialize.InitWebEngine()
	listenAddr := fmt.Sprintf("%s:%d", config.Config.ListenAddr, config.Config.ListenPort)

	logger.Info("Starting the web service engine, listening on ", listenAddr)
	err = eng.Run(listenAddr)
	if err != nil {
		logger.Error("Startup failed, error in starting the Web service engine: ", err)
		return
	}
}
