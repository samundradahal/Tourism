package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"root": "root@123",
	})
}
