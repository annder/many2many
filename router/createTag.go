package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"many2many/dto"
	"many2many/global"
	"many2many/model"
)

func CreateTag(context *gin.Context) {
	tagDTO := dto.Tag{}
	t := model.Tag{}
	if err := context.ShouldBindWith(&tagDTO, binding.Form); err != nil {
		if errs, ok := err.(validator.ValidationErrors); !ok {
			context.JSON(200, gin.H{
				"data":    nil,
				"message": errs.Translate(global.Trans),
				"code":    302,
			})
			context.Abort()
			return
		}
	}
	if ok := t.Create(tagDTO.Name, tagDTO.Desc); ok {
		context.JSON(200, gin.H{
			"data":    ok,
			"message": "请求成功",
			"code":    200,
		})
		return
	}
}
