package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"many2many/db"
	"many2many/inits"
	"many2many/router"
)

func main() {
	inits.Table()
	if err := inits.ValidatorTrans("zh"); err != nil {
		fmt.Printf("inits trans failed, err:%v\n", err)
		return
	}
	s := gin.Default()
	s.GET("/", func(context *gin.Context) {
		context.String(200, "ok")
	})
	s.POST("/create_article", router.CreateArticle)
	s.POST("/create_tag", router.CreateTag)
	s.GET("/article_list",router.GetAllArticle)
	defer db.CloseMysql()
	log.Fatal(s.Run(":8080"))

}
