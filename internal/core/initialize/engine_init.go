package initialize

import (
	"github.com/gin-gonic/gin"
	"web-scaffold/internal/core/config"
	"web-scaffold/internal/router"
)

// @Author: yv1ing
// @Email:  me@yvling.cn
// @Date:   2025/10/28 11:27
// @Desc:	初始化Web服务引擎

func InitWebEngine() *gin.Engine {
	gin.SetMode(config.Config.Mode)

	eng := gin.New()
	router.InitRouter(eng)
	
	return eng
}
