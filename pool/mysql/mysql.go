package mysql

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	// set db dirver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// set default database
	err := orm.RegisterDataBase("default", "mysql", "gameva:v8hdDTJy3c2ri5YB@tcp(192.168.101.3:3306)/stelin?timeout=90s&charset=utf8", 30)
	if err != nil {
		fmt.Println("mysql connect error=", err)
	}
}
