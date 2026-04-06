package middlewares

import (
	"go-shopping/net"
	"go-shopping/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, net.NewRes(http.StatusUnauthorized, nil, "未提供认证令牌"))
			c.Abort()
			return
		}

		// 解析 Bearer token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, net.NewRes(http.StatusUnauthorized, nil, "认证令牌格式错误"))
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := utils.ParseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, net.NewRes(http.StatusUnauthorized, nil, "认证令牌无效"))
			c.Abort()
			return
		}

		// 从claims中提取用户ID
		mapClaims, ok := claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, net.NewRes(http.StatusUnauthorized, nil, "认证令牌格式错误"))
			c.Abort()
			return
		}

		userID, ok := mapClaims["sub"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, net.NewRes(http.StatusUnauthorized, nil, "认证令牌格式错误"))
			c.Abort()
			return
		}

		// 将用户ID存储到上下文中
		c.Set("userID", uint(userID))
		c.Next()
	}
}
