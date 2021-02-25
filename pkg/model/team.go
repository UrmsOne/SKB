/*
@Author: urmsone urmsone@163.com
@Date: 2/25/21 9:28 PM
@Name: team.go
*/
package model

import "time"

type Team struct {
	// gorm.Model
	ID           uint `gorm:"primary_key"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Avatar       string
	Name         string
	Position     string
	Introduction string // 简介

	LastHandlerTime time.Time
	State           uint // 0: 1: 2:
}
