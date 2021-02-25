/*
@Author: urmsone urmsone@163.com
@Date: 2/25/21 3:07 PM
@Name: gorm.go
*/
package main

import (
	"SKB/pkg/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func DbInit(username, password, host, port, name string) error {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", username, password, host, port, name)
	var err error
	db, err = gorm.Open("mysql", uri)
	if err != nil {
		return err
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return nil
}

func main() {
	if err := DbInit("root", "123456", "192.168.253.129", "3310", "server_demo"); err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//db.Debug().AutoMigrate(
	//	model.User{},
	//)
	//var u model.User
	//if err := db.First(&u); err != nil {
	//	//gorm.Model{}
	//	fmt.Println(err)
	//}
	//fmt.Println(u)
	//db.Create(&model.User{
	//	UserName: "urmsone",
	//	NickName: "urmsone",
	//})
	////fmt.Println(u)
	//db.First(&u)
	//fmt.Println(u)
	u := &model.User{NickName: "urmsone", UserName: "urmsone"}
	u1 := &model.User{}
	if err := db.Model(&u).First(&u1).Error; err != nil {
		fmt.Println(err)
		fmt.Println("user is exist")
	}
	db.Unscoped().Delete(&model.User{ID: 1})
	fmt.Println(u1)
}
