package model

import "many2many/db"

type Article struct {
	Model
	Title   string `json:"title"`   // 标题
	Content string `json:"content"` // 内容
	Tag     []Tag  `gorm:"many2many:article_tags"`
}

func (*Article) Create(title, content string, t []Tag) bool {
	mysql := db.Mysql().Debug()
	if err := mysql.Create(&Article{Title: title, Content: content, Tag: t}); err != nil {
		return false
	}
	return true
}
func (*Article) FindArticleByTag() []map[string]interface{} {
	a := []map[string]interface{}{{}}
	mysql := db.Mysql().Debug()
	mysql.Table("man_article as ma").
		Joins("LEFT JOIN man_article_tags AS mat ON ma.id = mat.article_id").
		Joins("LEFT JOIN man_tag AS mt ON mt.id = mat.tag_id").
		Select("ma.title,ma.content,mt.name").
		Scan(&a)

	return a
}

func (*Article) FindAll() ([]Article, error) {
	a := []Article{{}}
	if err := db.Mysql().Debug().Find(&a).Error; err != nil {
		return a, err
	}
	return a, nil
}

func (*Article) Find(title, content string) (Article, error) {
	var a Article
	if err := db.Mysql().Debug().Where("title = ? or content = ?", title, content).Find(&a).Error; err != nil {
		return a, err
	}
	return a, nil
}
