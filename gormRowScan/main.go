package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "test:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.LogMode(true)
	db.SingularTable(true)
	var id int64
	var name string
	db.Table("user").Where("name = ?", "jinzhu").Select("id ,name").Row().Scan(&id, &name)

	fmt.Println("id:", id)
	fmt.Println("name:", name)
}
