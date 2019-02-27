package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Address struct {
	Name     sql.NullString `gorm:"size:64;unique" json:"name"`
	Province string         `gorm:"size:64" json:"province"`
	City     string         `gorm:"size:64" json:"city"`
	District string         `gorm:"size:64" json:"district"`
}

type User struct {
	//gorm.Model
	Email                 string `form:"email"`
	Password              string
	Name                  string `form:"name"`
	Gender                string
	Role                  string
	Birthday              time.Time
	Balance               float32
	DefaultBillingAddress uint      `gorm:"-"`
	Addresses             []Address //无法识别,没有生成对应的列

}

func main() {
	db, err := gorm.Open("mysql", "test:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.LogMode(true)
	db.SingularTable(true)

	var (
		id   int
		name string
	)
	rows, err := db.DB().Query("select id, name from user where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
