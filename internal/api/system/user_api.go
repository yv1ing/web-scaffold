package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web-scaffold/internal/core/config"
	"web-scaffold/internal/core/constant"
	"web-scaffold/pkg/auth"
	"web-scaffold/pkg/encrypt"

	systemmodel "web-scaffold/internal/model/system"
	systemservice "web-scaffold/internal/service/system"
)

// @Author: yv1ing
// @Email:  me@yvling.cn
// @Date:   2025/10/28 15:09
// @Desc:	系统用户接口实现

// UserLoginHandler 系统用户登入
func UserLoginHandler(ctx *gin.Context) {
	type reqType struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req reqType
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, systemmodel.Response{
			Code: constant.API_INVALID_REQUEST_PARAMS,
			Info: "invalid request params",
		})
		return
	}

	user, err := systemservice.FindUserByUsername(req.Username)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, systemmodel.Response{
				Code: constant.API_PERMISSION_DENIED,
				Info: "username or password incorrect",
			})
			return
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, systemmodel.Response{
				Code: constant.API_FIND_DATA_ERROR,
				Info: "failed to find user",
			})
			return
		}
	} else {
		if user.Password != encrypt.Sha256String(req.Password, config.Config.SecretKey) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, systemmodel.Response{
				Code: constant.API_PERMISSION_DENIED,
				Info: "username or password incorrect",
			})
			return
		}
	}

	jwtSign := encrypt.RandomString(32)
	jwtToken, err := auth.CreateAccessToken(user.ID, user.Username, config.Config.SecretKey, jwtSign)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, systemmodel.Response{
			Code: constant.API_INTERNAL_ERROR,
			Info: "failed to generate access token",
		})
		return
	}

	user.JwtSign = jwtSign
	err = systemservice.UpdateUser(user.ID, "", "", "", "", "", "", "", jwtSign)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, systemmodel.Response{
			Code: constant.API_UPDATE_DATA_ERROR,
			Info: "failed to update user status",
		})
		return
	}

	ctx.JSON(http.StatusOK, systemmodel.Response{
		Code: constant.API_OPERATE_SUCCESS,
		Info: "login success",
		Data: gin.H{
			"jwt_token": jwtToken,
		},
	})
}

// UserLogoutHandler 系统用户登出
func UserLogoutHandler(ctx *gin.Context) {
	userID := ctx.MustGet("user_id")
	user, err := systemservice.FindUserByID(userID.(uint))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, systemmodel.Response{
			Code: constant.API_FIND_DATA_ERROR,
			Info: "failed to find user",
		})
		return
	}

	err = systemservice.UpdateUser(user.ID, "", "", "", "", "", "", "", "-")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, systemmodel.Response{
			Code: constant.API_UPDATE_DATA_ERROR,
			Info: "failed to update user status",
		})
		return
	}

	ctx.JSON(http.StatusOK, systemmodel.Response{
		Code: constant.API_OPERATE_SUCCESS,
		Info: "logout success",
	})
}

// CreateUserHandler 创建系统用户
func CreateUserHandler(ctx *gin.Context) {
	type reqType struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Avatar   string `json:"avatar"`
		Role     string `json:"role"`
	}

	var (
		req reqType
		err error
	)
	err = ctx.ShouldBindBodyWithJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, systemmodel.Response{
			Code: constant.API_INVALID_REQUEST_PARAMS,
			Info: "invalid request params",
		})
		return
	}

	err = systemservice.CreateUser(req.Username, req.Password, req.Name, req.Email, req.Phone, req.Avatar, req.Role)
	if err != nil {
		if err.Error() == "username already exists" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, systemmodel.Response{
				Code: constant.API_INVALID_REQUEST_PARAMS,
				Info: "username already exists",
			})
			return
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, systemmodel.Response{
				Code: constant.API_CREATE_DATA_ERROR,
				Info: "failed to create user",
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, systemmodel.Response{
		Code: constant.API_OPERATE_SUCCESS,
		Info: "create user success",
	})
}

// DeleteUserHandler 删除系统用户
func DeleteUserHandler(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, systemmodel.Response{
			Code: constant.API_INVALID_REQUEST_PARAMS,
			Info: "invalid request params",
		})
		return
	}

	err = systemservice.DeleteUser(uint(userID))
	if err != nil {
		if err.Error() == "record not found" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, systemmodel.Response{
				Code: constant.API_FIND_DATA_ERROR,
				Info: "record not found",
			})
			return
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, systemmodel.Response{
				Code: constant.API_DELETE_DATA_ERROR,
				Info: "failed to delete user",
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, systemmodel.Response{
		Code: constant.API_OPERATE_SUCCESS,
		Info: "delete user success",
	})
}

// UpdateUserHandler 更新系统用户
func UpdateUserHandler(ctx *gin.Context) {
	type reqType struct {
		UserID   uint   `json:"user_id"`
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Avatar   string `json:"avatar"`
		Role     string `json:"role"`
	}

	var (
		req reqType
		err error
	)

	err = ctx.ShouldBindBodyWithJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, systemmodel.Response{
			Code: constant.API_INVALID_REQUEST_PARAMS,
			Info: "invalid request params",
		})
		return
	}

	err = systemservice.UpdateUser(req.UserID, req.Username, req.Password, req.Name, req.Email, req.Phone, req.Avatar, req.Role, "")
	if err != nil {
		if err.Error() == "record not found" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, systemmodel.Response{
				Code: constant.API_FIND_DATA_ERROR,
				Info: "record not found",
			})
			return
		} else if err.Error() == "username already exists" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, systemmodel.Response{
				Code: constant.API_INVALID_REQUEST_PARAMS,
				Info: "username already exists",
			})
			return
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, systemmodel.Response{
				Code: constant.API_UPDATE_DATA_ERROR,
				Info: "failed to update user",
			})
			return
		}
	}

	ctx.JSON(http.StatusOK, systemmodel.Response{
		Code: constant.API_OPERATE_SUCCESS,
		Info: "update user success",
	})
}

// FindUserHandler 条件查询系统用户
func FindUserHandler(ctx *gin.Context) {
	switch ctx.Query("type") {
	case "user_id":
		userID, err := strconv.Atoi(ctx.Query("user_id"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, systemmodel.Response{
				Code: constant.API_INVALID_REQUEST_PARAMS,
				Info: "invalid request params",
			})
			return
		}
		user, err := systemservice.FindUserByID(uint(userID))
		if err != nil {
			if err.Error() == "record not found" {
				ctx.AbortWithStatusJSON(http.StatusNotFound, systemmodel.Response{
					Code: constant.API_FIND_DATA_ERROR,
					Info: "record not found",
				})
				return
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, systemmodel.Response{
					Code: constant.API_FIND_DATA_ERROR,
					Info: "failed to find user",
				})
				return
			}
		}

		ctx.JSON(http.StatusOK, systemmodel.Response{
			Code: constant.API_OPERATE_SUCCESS,
			Info: "find user success",
			Data: gin.H{
				"total": 1,
				"users": []systemmodel.User{*user},
			},
		})
		break

	case "username":
		username := ctx.Query("username")
		user, err := systemservice.FindUserByUsername(username)
		if err != nil {
			if err.Error() == "record not found" {
				ctx.AbortWithStatusJSON(http.StatusNotFound, systemmodel.Response{
					Code: constant.API_FIND_DATA_ERROR,
					Info: "record not found",
				})
				return
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, systemmodel.Response{
					Code: constant.API_FIND_DATA_ERROR,
					Info: "failed to find user",
				})
				return
			}
		}

		ctx.JSON(http.StatusOK, systemmodel.Response{
			Code: constant.API_OPERATE_SUCCESS,
			Info: "find user success",
			Data: gin.H{
				"total": 1,
				"users": []systemmodel.User{*user},
			},
		})
		break

	case "name":
		name := ctx.Query("name")
		users, err := systemservice.FindUserByName(name)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, systemmodel.Response{
				Code: constant.API_FIND_DATA_ERROR,
				Info: "failed to find user",
			})
			return
		}
		if len(users) == 0 {
			ctx.AbortWithStatusJSON(http.StatusNotFound, systemmodel.Response{
				Code: constant.API_FIND_DATA_ERROR,
				Info: "record not found",
			})
			return
		}

		ctx.JSON(http.StatusOK, systemmodel.Response{
			Code: constant.API_OPERATE_SUCCESS,
			Info: "find user success",
			Data: gin.H{
				"total": len(users),
				"users": users,
			},
		})
		break

	default:
		ctx.AbortWithStatusJSON(http.StatusBadRequest, systemmodel.Response{
			Code: constant.API_INVALID_REQUEST_PARAMS,
			Info: "invalid request params",
		})
		return
	}
}

// ListUserHandler 分页查询系统用户列表
func ListUserHandler(ctx *gin.Context) {
	_page := ctx.DefaultQuery("page", "1")
	_size := ctx.DefaultQuery("size", "10")

	page, err := strconv.Atoi(_page)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, systemmodel.Response{
			Code: constant.API_INVALID_REQUEST_PARAMS,
			Info: "invalid request params",
		})
		return
	}
	size, err := strconv.Atoi(_size)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, systemmodel.Response{
			Code: constant.API_INVALID_REQUEST_PARAMS,
			Info: "invalid request params",
		})
		return
	}

	users, total, err := systemservice.FindUserListWithPage(page, size)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, systemmodel.Response{
			Code: constant.API_FIND_DATA_ERROR,
			Info: "failed to find user",
		})
		return
	}

	ctx.JSON(http.StatusOK, systemmodel.Response{
		Code: constant.API_OPERATE_SUCCESS,
		Info: "find user success",
		Data: gin.H{
			"users": users,
			"total": total,
		},
	})
}
