package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
type Region struct {
	Id       int
	Name     string
	ParentId int
}
type User struct {
	gorm.Model
	Name     string	`gorm:"size:255"`
	Email    string	`gorm:"size:255"`
	Password string	`gorm:"size:255"`
}
func main() {

	db, err := gorm.Open("mysql", "root:root@/go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("connetion err %s", err)
	}
	defer db.Close()
	//tableExist(db)
	createUser(db)
	u := &User{}
	u.Name = "wanna"
	u.Email = "2852@qq.com"
	u.Password = "password"
	db.Create(u)
}

func tableExist(db *gorm.DB) {
	if !db.HasTable(&Region{}) {
		fmt.Println("Regions 表不存在！\n")
	}
}
func createUser(db *gorm.DB) {
	db.DropTableIfExists(&User{})

	db.CreateTable(&User{})
	//db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
	//fmt.Printf("%s", table)
}
// 加这个可以重命名表名，默认用的user的复数
func (u *User) TableName() string {
	return "admin-users"
}

