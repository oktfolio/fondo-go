package middleware

import (
	"fondo-go/auth"
	"fondo-go/casbin"
	"fondo-go/common"
	"fondo-go/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

const AuthHeaderField = "Authorization"

func parseClaims(c *gin.Context) *auth.Claims {
	atoken := c.GetHeader(AuthHeaderField)

	if atoken == "" {
		atoken = c.Query("token")

	}

	if strings.TrimSpace(atoken) == "" {
		return nil
	}

	return auth.ParseToken(atoken, config.Conf.Env.Token.SecretKey)
}

// Auth 验证中间件
func Auth(isClient bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := parseClaims(c)
		if claims == nil {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}
		if isClient {
			if claims.User == nil {
				c.JSON(http.StatusUnauthorized, nil)
				c.Abort()
				return
			}
			c.Set("user", claims.User)
		}
		c.Next()
	}
}

func CasbinAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := parseClaims(c)
		user := claims.User.Username
		path := c.Request.URL.Path
		method := c.Request.Method

		if !casbin.E.Enforce(user, path, method) {

			result := common.Result{
				Code:      20000,
				Message:   "login success",
				Timestamp: time.Now().Unix(),
			}

			c.JSON(http.StatusUnauthorized, result)
			c.Abort()
			return
		}
		c.Next()
	}
}
