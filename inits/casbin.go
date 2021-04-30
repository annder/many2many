package inits

import (
	"log"
	"many2many/global"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

// 域的RBAC
func RBACInit() {
	a, _ := gormadapter.NewAdapter("mysql", "root:1234567@tcp(127.0.0.1:3306)/many2many?charset=utf8mb4&parseTime=True&loc=Local")
	global.Enforcer, _ = casbin.NewEnforcer("conf/model.conf", a)
	err := global.Enforcer.LoadPolicy()
	if err != nil {
		log.Fatal(err)
	}
	err = global.Enforcer.SavePolicy()
	if err != nil {
		log.Fatalln(err)
	}
}
