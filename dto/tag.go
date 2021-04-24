package dto

type Tag struct {
	Name string `form:"name" json:"name" binding:"required"`
	Desc string `fom:"name" json:"desc"`
}
