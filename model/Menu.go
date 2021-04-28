package model

type Menu struct {
	MenuID     int      `json:"menu_id"`
	RoleID     int      `json:"role_id"`
	Name       string   `json:"name" gorm:"index,unique;not null"` // 菜单名
	RoutePath  string   `json:"route_path" gorm:"unique;not null"` // 路径
	Icon       string   `json:"icon"`
	Hide       bool     `json:"hide" gorm:"unique;not null"`
	Component  string   `json:"component"`          // 组件， 可能是Layout文件
	Hidden     bool     `json:"hidden"`             // 隐藏 是否隐藏
	AlwaysShow bool     `json:"always_show"`        // 永久展现
	Redirect   string   `json:"redirect"`           // 重定向
	Mate       MenuMate `json:"menu_mate" gorm:"-"` // 这个为啥有问题？？
	MenuMateID int      `json:"menu_mate_id"`
	Children   []Menu   `json:"children" gorm:"-"` // 子路径
}
