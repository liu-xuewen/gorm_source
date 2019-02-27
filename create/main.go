package main

import (
	"database/sql"
	"fmt"
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
	//	gorm.Model
	ID                    uint `gorm:"primary_key"`
	CreatedAt             int64
	UpdatedAt             int64
	DeletedAt             *int64 `sql:"index"`
	Email                 string `form:"email"`
	Password              string
	Name                  string `form:"name"`
	Gender                string //`gorm:"column:"`
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
	//SetDirectPanicWhenErrCreate(db)
	user := User{Name: "Jinzhu", Email: "123456@qq.com", Birthday: time.Now()}

	//将user作为scope的值，创建一个scope.
	//isBefore := db.NewRecord(user) // => returns `true` as primary key is blank
	///fmt.Println("is_before是否存在", isBefore)
	if err := db.Create(&user).Error; err != nil {
		fmt.Println("create user error:", err)
	}
	//fmt.Println("user----", user)
	if err := db.Delete(&user).Error; err != nil {
		fmt.Println("create user error:", err)
	}
	//isAfter := db.NewRecord(user) // => return `false` after `user` created
	///fmt.Println("is_after是否存在", isAfter)
}

/*

// Rollback the insertion if user's id greater than 1000
func (u *User) AfterSave() (err error) {
	fmt.Println("1111111")
	//debug.PrintStack()
	if u.Name == "Jinzhu" {
		err = errors.New("user name is jinzhu")
	}
	return
}

// Rollback the insertion if user's id greater than 1000
func (u *User) BeforeCreate() (err error) {
	fmt.Println("11666666666661")
	//debug.PrintStack()
	if u.Name == "Jinzhu" {
		err = errors.New("user name is jinzhu")
	}
	return
}

func SetDirectPanicWhenErrCreate(db *gorm.DB) {
	fmt.Println("22222")
	db.Callback().Create().After("gorm:create").Register("plugin:after_create_err_panic", func(scope *gorm.Scope) {
		if scope.HasError() {
			panic(scope.DB().Error.Error())
		}
	})
}
*/
