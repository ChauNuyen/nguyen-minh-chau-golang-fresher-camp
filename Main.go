package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type user struct {
	Id      int    `json:"id, omitempty" gorm:"column:id;"`
	Address string `json:"country_code" gorm:"column:address;"`
	Phone   string `json:"district" gorm:"column:phone;"`
}

func (user) TableUser() string {
	return "users"
}

func main() {
	dsn := os.Getenv("DBConnectionStr_user")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	//insert new user
	newUser := user{Address: "A", Phone: "000"}
	if err := db.Create(&newUser); err != nil {
		fmt.Print(newUser)
	}

	//Select user from table
	var users []user
	if err := db.Where("address", "none").Find(&users); err != nil {
		log.Println(err)
	}
	fmt.Println(users)

	//delete user
	db.Table(user{}.TableUser()).Where("id = 11").Delete(nil)
	fmt.Println(users)

	//update user
	db.Table(user{}.TableUser()).Where("id = 12").Updates(map[string]interface{}{
		"Address": "VN",
	})
	fmt.Println(users)

}
