package router

import (
	"github.com/gin-gonic/gin"

	systemapi "web-scaffold/internal/api/system"
)

// @Author: yv1ing
// @Email:  me@yvling.cn
// @Date:   2025/10/28 15:07
// @Desc:	初始化Web应用路由

func InitRouter(eng *gin.Engine) {
	api := eng.Group("/api")

	// 系统内置路由
	sys := api.Group("/sys")

	sys.GET("/users/list", systemapi.ListUserHandler)
	sys.GET("/users/find", systemapi.FindUserHandler)
	sys.POST("/users/create", systemapi.CreateUserHandler)
	sys.DELETE("/users/delete", systemapi.DeleteUserHandler)
	sys.PUT("/users/update", systemapi.UpdateUserHandler)

	// 实际业务路由
}
