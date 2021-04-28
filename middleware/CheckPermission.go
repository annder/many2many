package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"many2many/global"
)

func notPermissionRequestResponse(c *gin.Context) {
	c.AbortWithStatusJSON(403, map[string]interface{}{
		"message": "没有权限访问此接口",
		"code":    403,
	})
}

func CheckPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.MustGet("role").(string)
		if r == "admin" {
			c.Next()
		}
		m := c.Request.Method
		p := c.Request.URL.Path
		log.Println("["+m+"]：", p)
		ok, err := global.Enforcer.Enforce(r, p, m)
		if err != nil {
			notPermissionRequestResponse(c)
		}
		if !ok {
			notPermissionRequestResponse(c)
		}
		c.Next()
	}
}
