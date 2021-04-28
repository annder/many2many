package db

import "gorm.io/gorm"

func FindByUserID(role string, userid int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if role == "admin" {
			return db
		}
		return db.Where("user_id = ?", userid)
	}
}
