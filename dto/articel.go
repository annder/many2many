package dto

type Article struct {
	Content string   `form:"content" json:"content" binding:"required"`
	Title   string   `form:"title" json:"title" binding:"required"`
	Tags    []string `form:"tags[]" json:"tags[]" binding:"required"`
}
