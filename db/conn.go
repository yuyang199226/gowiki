package db

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
	"gowiki/config"
)

var MysqlDB *sql.DB


func initEngine() error {
	var err error
	config := config.Get()
	dsn := fmt.Sprintf("%s:%s@/%s", config.DataBase.User, config.DataBase.Password, config.DataBase.Name)
	MysqlDB, err = sql.Open("mysql", dsn)
	MysqlDB.SetMaxIdleConns(2)
	MysqlDB.SetMaxOpenConns(10)
	if err != nil {
		return err
	}
	err = MysqlDB.Ping()
	return err
}

func init() {
	if err := initEngine(); err != nil {
		panic(err)
	}

}

