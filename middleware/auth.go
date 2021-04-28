package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"many2many/db"
	"many2many/model"
)

var user model.User

func Auth() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(model.User); ok {
				return jwt.MapClaims{
					"user_id": v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(context *gin.Context) interface{} {
			u := model.User{}
			if err := db.Mysql().Where("id = ?", jwt.ExtractClaims(context)).Find(&u); err != nil {
				context.AbortWithStatusJSON(200, map[string]interface{}{
					"code":    500,
					"message": err.Error,
					"data":    nil,
				})
			}
			return u
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			username := c.DefaultQuery("username", "")
			password := c.DefaultQuery("password", "")
			u := model.User{}
			if err := db.Mysql().Where("username = ? and password = ?", username, password).Find(&u).Error; err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return u, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(model.User); ok {
				if err := db.Mysql().Where("id = ?", v.ID); err != nil {
					return false
				}
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	return authMiddleware
}
