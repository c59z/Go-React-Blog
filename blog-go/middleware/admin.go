package middleware

import (
	"blog-go/model/appTypes"
	"blog-go/model/response"
	"blog-go/utils"

	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID := utils.GetRoleID(c)
		if roleID != appTypes.Admin {
			response.Forbidden("Access denied. Admin privileges are required", c)
			c.Abort()
			return
		}

		c.Next()
	}
}
