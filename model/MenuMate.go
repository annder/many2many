package model

import "gorm.io/gorm"

type Mate struct {
	gorm.Model
	Title      string `json:"title"`
	Icon       string `json:"icon"`
	NoCache    bool   `json:"no_cache"`   // 没有缓存
	Breadcrumb bool   `json:"breadcrumb"` // 面包屑
	Affix      bool   `json:"affix"`
	MenuID     int    `json:"menu_id"`
}
