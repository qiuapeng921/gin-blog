package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v2"
	"github.com/jinzhu/gorm"
)

var Adapter *gormadapter.Adapter
var Enforcer *casbin.Enforcer

func AuthInit(client *gorm.DB) {
	var err error
	// 将数据库连接同步给插件， 插件用来操作数据库
	Adapter, err = gormadapter.NewAdapterByDB(client)
	if err != nil {
		panic("链接数据库失败" + err.Error())
		return
	}
	// 这里也可以使用原生字符串方式
	Enforcer, err = casbin.NewEnforcer("auth_model.conf", Adapter)
	if err != nil {
		panic("加载配置失败" + err.Error())
	}
	// 开启权限认证日志
	Enforcer.EnableLog(true)
	// 加载数据库中的策略
	err = Enforcer.LoadPolicy()
	if err != nil {
		fmt.Println("loadPolicy error")
		return
	}
	// 创建一个角色,并赋于权限
	// admin 这个角色可以访问GET 方式访问 /api/v2/ping
	res, err := Enforcer.AddPolicy("admin", "/api/v2/ping", "GET")
	if !res {
		fmt.Println("policy is exist")
	} else {
		fmt.Println("policy is not exist, adding")
	}
	// 将 test 用户加入一个角色中
	_, _ = Enforcer.AddRoleForUser("test", "root")
	_, _ = Enforcer.AddRoleForUser("admin", "admin")
	// 请看规则中如果用户名为 root 则不受限制
}
