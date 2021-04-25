package model

type Api struct {
	Name   string `json:"name"` // API名
	RoleID int    `json:"role_id"`
	Method string `json:"method"` // 方法
	Url    string `json:"url"`    // url名
	Group  string `json:"group"`  // API 组名
}
