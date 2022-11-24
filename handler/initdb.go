package handler

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"strings"
)

var DB *sql.DB

func InitDB() *sql.DB {
	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config := viper.New()

	config.AddConfigPath(path)     //设置读取的文件路径
	config.SetConfigName("config") //设置读取的文件名
	config.SetConfigType("yaml")   //设置文件的类型
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	username:=fmt.Sprintf("%v",config.Get("db.username"))
	password:=fmt.Sprintf("%v",config.Get("db.password"))
	host:=fmt.Sprintf("%v",config.Get("db.host"))
	port:=fmt.Sprintf("%v",config.Get("db.port"))
	dbname:=fmt.Sprintf("%v",config.Get("db.dbname"))

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