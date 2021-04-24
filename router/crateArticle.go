package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"many2many/dto"
	"many2many/global"
	"many2many/model"
)

func CreateArticle(context *gin.Context) {
	article := dto.Article{}
	t := model.Tag{}
	a := model.Article{}
	if err := context.ShouldBindWith(&article, binding.Form); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			context.JSON(200, gin.H{
				"message": errs.Translate(global.Trans),
				"data":    nil,
				"code":    301,
			})
			context.Abort()
			return
		}
	}
	tag := article.Tags
	tagResult, _ := t.Find(tag)
	if tagResult == nil {
		if ok := a.Create(article.Title, article.Content, tagResult); ok {
			context.JSON(200, gin.H{
				"message": "创建成功",
				"data":    ok,
				"code":    200,
			})
			context.Abort()
			return
		}
		// 创建一个Tag
	} else {
		for i := 1; i < len(article.Tags); i++ {
			t.Create(article.Tags[i], "")
		}
		// 还要找出tag
		ts, _ := t.Find(article.Tags)
		if ok := a.Create(article.Title, article.Content, ts); ok {
			context.JSON(200, gin.H{
				"message": "创建失败,服务器内部错误",
				"data":    !ok,
				"code":    500,
			})
			context.Abort()
			return
		} else {
			context.JSON(200, gin.H{
				"message": "创建成功",
				"data":    ok,
				"code":    200,
			})
			return
		}
	}
}

func GetAllArticle (ctx *gin.Context){
	a := model.Article{}
	ctx.JSON(200,a.FindArticleByTag())
	return
}