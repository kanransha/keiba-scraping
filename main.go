package main

import (
	"os"
	"fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	USER := "root"
	PASS := os.Getenv("MYSQL_ROOT_PASSWORD")
	HOST := "database"
	DBNAME := "keiba"
	CONNECT := USER + ":" + PASS + "@tcp(" + HOST + ")" + "/" + DBNAME
	db, err := gorm.Open("mysql", CONNECT + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
    panic(err.Error())
	}
	fmt.Print("接続に成功しました！")
  defer db.Close()
}