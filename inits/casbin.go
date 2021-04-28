package inits

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"log"
	"many2many/db"
	"many2many/global"
	"many2many/model"
)

// 域的RBAC
func RBACInit() {
	casbinRule := model.CasbinRule{}
	a, _ := gormadapter.NewAdapterByDBWithCustomTable(db.Mysql(), &casbinRule)
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
