/*
@Author: urmsone urmsone@163.com
@Date: 2/25/21 9:27 PM
@Name: advertising.go
@Description: 广告表
*/
package model

import "time"

type Advertising struct {
	// gorm.Model
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Img       string
	Name      string
	Link      string
	State     uint // 0: 1: 2:
	Remarks   string
}
