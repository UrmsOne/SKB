/*
@Author: urmsone urmsone@163.com
@Date: 2/25/21 9:28 PM
@Name: order.go
*/
package model

import "time"

type Order struct {
	// gorm.Model
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt *time.Time `sql:"index"` // 启动软删除,这个字段必须是指针类型,否则db.First/db.Find的时候查不出结果
	// attribute
	Account         string
	UserName        string
	KeyImg          string    // 主图
	Theme           string    // 主题
	Price           float64   // 单价
	Amount          int       // 数量
	DiscountedPrice float64   // 优惠金额
	RealPaid        float64   // 实付金额
	OrderTime       time.Time // 下单时间
	State           uint      // 状态 0: 1: 2:
	// foreign keys
}
