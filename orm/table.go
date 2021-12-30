package orm

import (
	"encoding/json"
	"fmt"
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

func InitTable(db *gorm.DB) {

	if db == nil{
		fmt.Println("Table Migrate error")
	}
	db.AutoMigrate(&User{})

}

func CreateData(db *gorm.DB){
	if db == nil{
		fmt.Println("Table CreateData error")
	}
	var user = []User{
		{Id: 1,
			Name:     "zs",
			Age:      18,
			Birthday: nil,
			Email:    "123@qq.com",
			PassWord: "123",
		},
		{Id: 2,
			Name:     "ls",
			Age:      13,
			Birthday: nil,
			Email:    "123@qq.com",
			PassWord: "123",
		},
		{Id: 3,
			Name:     "wr",
			Age:      14,
			Birthday: nil,
			Email:    "123@qq.com",
			PassWord: "123",
		},
		{Id: 4,
			Name:     "zl",
			Age:      17,
			Birthday: nil,
			Email:    "123@qq.com",
			PassWord: "123",
		},
		{Id: 20,
			Name:     "qq",
			Age:      18,
			Birthday: nil,
			Email:    "123@qq.com",
			PassWord: "123",
		},
	}
	db.Create(&user)


}


func SearchData(db *gorm.DB, val1 uint)  string{
	if db == nil{
		fmt.Println("Table SearchData error")
	}
	var user User
	db.Select(val1).Where("id = ?",val1).Find(&user)
	obj, err := json.Marshal(user)
	if err != nil{
		err.Error()
	}
	fmt.Println(string(obj))
	return string(obj)

}

func DeleteData(db *gorm.DB, val1 uint){
	if db == nil{
		fmt.Println("Table DeleteData error")
	}
	var user User
	user.Id = val1
	db.Debug().Delete(&user)

}

func UpdateData(db *gorm.DB, val1 uint){
	if db == nil{
		fmt.Println("Table UpdateData error")
	}
	var user User
	user.Id = val1
	db.Model(&user).Update("pass_word", "456")
}

func ChangePasswordFromId(db *gorm.DB, val1 uint, ret string){
	if db == nil{
		fmt.Println("Table UpdateData error")
	}
	var user User
	user.Id = val1
	db.Model(&user).Update("pass_word", ret)
}