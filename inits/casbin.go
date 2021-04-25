package inits

import (
	"github.com/casbin/casbin/v2"
	"log"
)

// 域的RBAC
func RBAC(){
	e,_ := casbin.NewEnforcer("conf/model.conf", "conf/policy.csv")
	log.Fatal(e.LoadPolicy())
}
