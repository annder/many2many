package model

type User struct {
	Model
	RoleID   int    `json:"role_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     []Role `json:"role" gorm:"many2many:user_role"`
}
