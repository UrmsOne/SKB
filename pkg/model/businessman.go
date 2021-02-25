/*
@Author: urmsone urmsone@163.com
@Date: 2/25/21 9:31 PM
@Name: businessman.go
*/
package model

import "time"

type Businessman struct {
	// gorm.Model
	ID            uint `gorm:"primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Img           string
	Name          string    // 商家名称
	Username      string    // 用户名称
	Phone         string    // 手机
	CommitTimes   int       // 提交次数
	CommittedTime time.Time // 提交时间
	ReviewTime    time.Time // 审核时间
	State         uint      // 状态 0: 1: 2:
}
