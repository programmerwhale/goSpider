package model

import (
	"fmt"
	"github.com/spf13/viper"
)

/*type Database struct {
	Self   *gorm.DB
}

var DB *sql.DB

func InitDB() {
	fmt.Println(viper.GetString("db.username"))
/*	path := strings.Join([]string{viper.GetString("db.username"), ":", PASSWORD, "@tcp(", HOST, ":", PORT, ")/", DBNAME, "?charset=utf8"}, "")
	fmt.Println(path)
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(10)
	DB.SetMaxIdleConns(5)
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")*/

}*/

func TestViper() {
	fmt.Println(viper.GetString("db.username"))
}