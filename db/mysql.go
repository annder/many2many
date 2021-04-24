package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func Mysql() *gorm.DB {
	dsn := "root:1234567@tcp(127.0.0.1:3306)/many2many?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: "man_", SingularTable: true}, // 单数表
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CloseMysql() error {
	db, err := Mysql().DB()
	if err != nil {
		log.Fatalln(err)
	}
	return db.Close()
}

