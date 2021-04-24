package model

import "many2many/db"

type Tag struct {
	Model
	Name    string    `json:"name" gorm:"unique;not null"` // 名字
	Desc    string    `json:"desc"`                        // 描述
	Article []Article `gorm:"many2many:article_tags"`      // 多对多
}

func (*Tag) Create(name, desc string) bool {
	if err := db.Mysql().Debug().Create(&Tag{
		Name: name,
		Desc: desc,
	}).Error; err != nil {
		return false
	}
	return true
}

func (*Tag) FindAll() ([]Tag, error) {
	t := []Tag{{}}
	mysql := db.Mysql().Debug()
	if err := mysql.Find(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}



func (*Tag) Find(name []string) ([]Tag, error) {
	t := []Tag{{}}
	mysql := db.Mysql().Debug()
	if err := mysql.Where("name in ?", name).Find(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}
