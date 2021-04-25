package model

type Menu struct {
	RoleID    int    `json:"role_id"`
	Name      string `json:"name" gorm:"index,unique;not null"` // 菜单名
	RoutePath string `json:"route_path" gorm:"unique;not null"` // 路径
	Icon      string `json:"icon"`
	RouteName string `json:"route_name" gorm:"unique;not null"`
	Hide      bool   `json:"hide" gorm:"unique;not null"`
}
