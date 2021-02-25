/*
@Author: urmsone urmsone@163.com
@Date: 2/25/21 9:29 PM
@Name: booking.go
@Description: 预约表
*/
package model

import "time"

type Booking struct {
	// gorm.Model
	ID              uint `gorm:"primary_key"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Type            string
	Phone           string
	LastHandlerTime time.Time
	State           uint   // 0: 1: 2:
	Intention       string // 意向描述
}
