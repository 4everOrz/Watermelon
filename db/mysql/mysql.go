package mysql

import (
	"Watermelon/common/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Orm *gorm.DB

func init() {
	var err error
	dataSourceName := config.GetString("Mysql_User") + ":" + config.GetString("Mysql_Password") + "@tcp(" +
		config.GetString("Mysql_IP") + ":" + config.GetString("Mysql_Port") + ")/" +
		config.GetString("Mysql_DBName") + "?charset=utf8"
	Orm, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("mysql init failed!！error:" + err.Error())
		return
	}
	Orm.DB().SetMaxIdleConns(10)
	Orm.DB().SetMaxOpenConns(100)
	Orm.SingularTable(true) //全局禁用表名复数
	fmt.Println("mysql init successed! ")

}
