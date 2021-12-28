package main

import (
	"fmt"
	"github.com/driver/mysql"
	"gorm.io/gorm"
	"time"
)



type User struct {
	Id          uint        `gorm:"AUTO_INCREMENT"`
	Name        string      `gorm:"size:50"`
	Age         int         `gorm:"size:3"`
	Birthday    *time.Time
	Email       string      `gorm:"type:varchar(50);unique_index"`
	PassWord    string      `gorm:"type:varchar(25)"`
}


var db *gorm.DB
func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "lhq:0318lhq@tcp(127.0.0.1:3306)/tlbb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}


	// 自动迁移数据结构(table schema)
	// 注意:在gorm中，默认的表名都是结构体名称的复数形式，比如User结构体默认创建的表为users
	// db.SingularTable(true) 可以取消表名的复数形式，使得表名和结构体名称一致
	db.AutoMigrate(&User{})



	// 插入记录
	db.Create(&User{Name:"bgbiao",Age:18,Email:"bgbiao@bgbiao.top"})
	db.Create(&User{Name:"xxb",Age:18,Email:"xxb@bgbiao.top"})
	//var user User
	var users []User
	// 查看插入后的全部元素
	fmt.Printf("插入后元素:\n")
	db.Find(&users)
	fmt.Println(users)
}