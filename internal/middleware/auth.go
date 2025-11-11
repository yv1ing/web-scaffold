package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"regexp"
	"strings"
	"web-scaffold/internal/core/config"
	"web-scaffold/internal/core/constant"
	"web-scaffold/pkg/auth"

	systemmodel "web-scaffold/internal/model/system"
	systemservice "web-scaffold/internal/service/system"
)

// @Author: yv1ing
// @Email:  me@yvling.cn
// @Date:   2025/10/29 11:33
// @Desc:	鉴权中间件

func extractBearerToken(c *gin.Context) string {
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		return ""
	}

	parts := strings.SplitN(authorization, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || parts[1] == "" {
		return ""
	}

	return parts[1]
}

func JwtAuthMiddleware(whitelist []string) gin.HandlerFunc {
	var whitelistRegex []*regexp.Regexp
	for _, pattern := range whitelist {
		re, err := regexp.Compile(pattern)
		if err == nil {
			whitelistRegex = append(whitelistRegex, re)
		}
	}

	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		for _, re := range whitelistRegex {
			if re.MatchString(path) {
				ctx.Next()
				return
			}
		}

		tokenStr := extractBearerToken(ctx)
		if tokenStr == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, systemmodel.Response{
				Code: constant.API_INVALID_REQUEST_HEADER,
				Info: "invalid request header",
			})
			return
		}

		claims, err := auth.ParseAccessToken(tokenStr, config.Config.SecretKey)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, systemmodel.Response{
					Code: constant.API_PERMISSION_DENIED,
					Info: "token has expired",
				})
			} else {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, systemmodel.Response{
					Code: constant.API_PERMISSION_DENIED,
					Info: "token is not valid",
				})
			}
			return
		}

		user, err := systemservice.FindUserByUsername(claims.Username)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, systemmodel.Response{
				Code: constant.API_FIND_DATA_ERROR,
				Info: "failed to find user",
			})
			return
		}
		if claims.JwtSign != user.JwtSign {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, systemmodel.Response{
				Code: constant.API_PERMISSION_DENIED,
				Info: "token has expired",
			})
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Set("username", claims.Username)
		ctx.Next()
	}
}
