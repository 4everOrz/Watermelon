package db

import (
	"Watermelon/config"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Params map[string]interface{}

var (
	Orm *sql.DB
)

func init() {
	var err error
	conf := config.GetConf()
	Orm, err = sql.Open(conf.SqlType, conf.SqlUser+":"+conf.SqlPassword+"@tcp("+conf.SqlHost+":"+conf.SqlPort+")/"+conf.SqlDB+"?charset=utf8mb4")
	if err != nil {
		log.Println(err)
	}
	Orm.SetMaxIdleConns(20)
	Orm.SetMaxOpenConns(20)
	if err := Orm.Ping(); err != nil {
		log.Println(err)
	}
}
