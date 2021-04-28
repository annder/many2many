package main

import (
	"fmt"
	"log"
	"many2many/db"
	"many2many/inits"
	"many2many/middleware"
	"many2many/router"

	"github.com/gin-gonic/gin"
)

func main() {
	inits.Table()
	inits.RBACInit()
	if err := inits.ValidatorTrans("zh"); err != nil {
		fmt.Printf("inits trans failed, err:%v\n", err)
		return
	}
	s := gin.Default()
	s.Use(middleware.LogerMiddleware())
	s.Use(middleware.CheckPermission()) // 检查权限
	s.GET("/", func(context *gin.Context) {
		context.String(200, "ok")
	})
	s.GET("/add_role", func(context *gin.Context) {

	})
	s.GET("/add_user_to_role", func(context *gin.Context) {

	})
	s.GET("/get_user", func(context *gin.Context) {

	})
	s.POST("/del_user", func(context *gin.Context) {

	})
	s.POST("/update_user", func(context *gin.Context) {

	})
	s.GET("/get_role", func(context *gin.Context) {

	})
	s.POST("/del_role", func(context *gin.Context) {

	})
	s.POST("/update_role", func(context *gin.Context) {

	})
	s.GET("/get_api", func(context *gin.Context) {

	})
	s.POST("/update_api", func(context *gin.Context) {

	})

	s.GET("/get_menu", func(context *gin.Context) {

	})
	// 获取我的目录
	s.POST("/update_menu", func(context *gin.Context) {

	})
	s.POST("/create_article", router.CreateArticle)
	s.POST("/create_tag", router.CreateTag)
	s.GET("/find_article_by_tag", router.GetAllArticle)
	defer db.CloseMysql()
	log.Fatal(s.Run(":8080"))

}
