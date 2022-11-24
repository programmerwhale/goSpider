package handler

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
)

var DB *sql.DB

func InitDB() *sql.DB {
	Config:=ViperHelper()
	username:=fmt.Sprintf("%v",Config.Get("db.username"))
	password:=fmt.Sprintf("%v",Config.Get("db.password"))
	host:=fmt.Sprintf("%v",Config.Get("db.host"))
	port:=fmt.Sprintf("%v",Config.Get("db.port"))
	dbname:=fmt.Sprintf("%v",Config.Get("db.dbname"))

	connectionString := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", dbname, "?charset=utf8"}, "")
	fmt.Println(connectionString)
	DB, _ = sql.Open("mysql", connectionString)
	DB.SetConnMaxLifetime(10)
	DB.SetMaxIdleConns(5)
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return nil
	}
	fmt.Println("connnect success")
	return DB
}