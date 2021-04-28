package global

import (
	"github.com/casbin/casbin/v2"
	ut "github.com/go-playground/universal-translator"
)

var (
	Trans    ut.Translator    // 全局验证器
	Enforcer *casbin.Enforcer // 全局casbin
)
