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
type WhiteDeviceId struct {
	Id int
	DeviceId string `gorm:"column:deviceId"`
}
type User struct {
	//gorm.Model
	Id uint32
	Username     string	`gorm:"size:255"`
	Email    string	`gorm:"size:255"`
	Password string	`gorm:"size:255"`
	Cellphone string `gorm:"column:cellphone"`
}

type DoctorUser struct {
	Name string
}
func main() {
	var err error
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/med_primary")
	if err != nil {
		fmt.Printf("connetion err %s", err)
	}
	defer db.Close()
	findByCellphone()
	//tableExist(db)
	//createUser(db)
	//data := &User{
	//	Id: 78389767,
	//	Username: "wanna",
	//}
	//err = db.Find(data).Error
	//var c,c1 int
	//err = db.Model(&User{}).Where("id", "0").Count(&c).Where("type","9").Count(&c1).Error
	//fmt.Println("qurey  %d, %d",c,c1)
	//if err != nil {
	//	fmt.Printf("err (%+v)", err)
	//	//fmt.Printf("data (%+v)", data)
	//}
	//u := &User{}
	//u.Name = "wanna"
	//u.Email = "2852@qq.com"
	//u.Password = "password"
	//db.Create(u)
}

func CreateByMap()  {
	 data :=  map[string]interface{}{
		"deviceId": "12345",
	}
	err := db.Model(&WhiteDeviceId{}).Save(data).Error
	if err != nil {
		println("save err: ",err)
	}
}

func findByCellphone()  {
	println("====")
	data := &User{
		Cellphone: "18227590006",
	}
	err := db.Where(data).First(data).Error
	if err != nil {
		fmt.Printf("err %+v",err)
	}
	fmt.Printf("%+v",data)
}

func CreateByStruct()  {
	data := &WhiteDeviceId{
		DeviceId: "12345677",
	}
	err := db.Save(data).Error
	fmt.Printf("%+v",data)
	if err != nil {
		println("save err: ",err)
	}
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
	return "user"
}

func (u *WhiteDeviceId) TableName() string {
	return "whiteDeviceId"
}

