package main

import (
	"database/sql"
	"errors"
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
	ID                    uint `gorm:"primary_key"`
	CreatedAt             int64
	UpdatedAt             int64
	DeletedAt             *int64 `sql:"index"`
	Email                 string `form:"email"`
	Password              string
	Name                  string `form:"name"`
	Gender                string
	Role                  string
	Birthday              *time.Time
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
	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	// Migrate the schema
	db.AutoMigrate(&User{})
}

// Rollback the insertion if user's id greater than 1000
func (u *User) AfterCreate() (err error) {
	if u.ID > 3 {
		err = errors.New("user id is already greater than 3")
	}
	return
}
