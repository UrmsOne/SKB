/*
@Author: urmsone urmsone@163.com
@Date: 2/25/21 9:31 PM
@Name: decoration.go
@Description: 装修设置表
*/
package model

import "time"

type Decoration struct {
	// gorm.Model
	ID              uint `gorm:"primary_key"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	KeyImg          string
	Carousel        string // 轮播图
	DetailImg       string // 详情页图
	Style           string // 装修风格
	Theme           string // 主题
	SubTheme        string // 副主题
	Price           float64 // 价格
	LastHandlerTime time.Time // 操作时间
	State           uint   // 状态 0: 1: 2:
}
