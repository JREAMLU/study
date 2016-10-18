package core

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var X *gorm.DB

func InitGorm(driver string) {
	var err error
	//"root:root@tcp(localhost:3306)/plucron?charset=utf8&parseTime=True&loc=Local"
	X, err = gorm.Open("mysql", driver)
	if err != nil {
		fmt.Println("err: ", err)
	}
	X.SingularTable(true)
	X.LogMode(true)
}
