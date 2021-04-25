package model

type Role struct {
	Model
	UserID int `json:"user_id"`
	Name string `json:"name"` // 角色名
	Desc string `json:"desc"` // 描述
	Menu []Menu `json:"menu"`
	Api  []Api  `json:"api"`
	User []User `json:"user" gorm:"many2many:role_user"`
}
