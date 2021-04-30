package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	MenuID     *int   `json:"menu_id"`
	RoleID     int    `json:"role_id"`
	MateID     int    `json:"mate_id"`
	Name       string `json:"name" gorm:"index,unique;not null"` // 菜单名
	RoutePath  string `json:"route_path" gorm:"unique;not null"` // 路径
	Icon       string `json:"icon"`
	Hide       bool   `json:"hide" gorm:"unique;not null"`
	Component  string `json:"component"`   // 组件， 可能是Layout文件
	Hidden     bool   `json:"hidden"`      // 隐藏 是否隐藏
	AlwaysShow bool   `json:"always_show"` // 永久展现
	Redirect   string `json:"redirect"`    // 重定向
	Mate       Mate
	Children   []Menu `json:"children" gorm:"foreginKey:MenuID"` // 子路径
}
