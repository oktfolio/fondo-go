package casbin

import (
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql"
)

var E *casbin.Enforcer

func InitCasbin()  {
	a := gormadapter.NewAdapter("mysql",
		"root:root@tcp(192.168.50.127:3306)/fondo-security?charset=utf8",
		true) // Your driver and data source.
	E = casbin.NewEnforcer("authz_model.conf", a)
}