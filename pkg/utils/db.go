/*
@Author: urmsone urmsone@163.com
@Date: 2/23/21 9:11 PM
@Name: db.go
*/
package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

func DbUtils() *gorm.DB {
	return db
}

func DbClose() error {
	if db != nil {
		if err := db.Close(); err != nil {
			return err
		}
	}
	return nil
}
