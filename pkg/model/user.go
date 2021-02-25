/*
@Author: urmsone urmsone@163.com
@Date: 2/24/21 2:00 PM
@Name: user.go
*/
package model

import (
	"SKB/pkg/utils"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	// gorm.Model
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt *time.Time `sql:"index"` // 启动软删除,这个字段必须是指针类型,否则db.First/db.Find的时候查不出结果
	// attribute
	Account      string
	NickName     string `gorm:"not null;unique_index"`
	UserName     string
	Avatar       string
	VipLevel     string
	Income       string
	ReferralCode string
	ReferOrder   string
	Client       string
	Consumption  string
	PasswordHash string
	Email        string
	// foreign keys
}

func GetUserByID(id string) (*User, error) {
	u := &User{}
	if err := utils.DbUtils().First(&u, id).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func GetUsers() ([]User, error) {
	var us []User
	if err := utils.DbUtils().Find(&us).Error; err != nil {
		return nil, err
	}
	return us, nil
}

func GetUser(u User) error {
	fmt.Println("model: ", u)
	var u1 []User
	if err := utils.DbUtils().Where(u).First(u1).Error; err != nil {
		fmt.Println("u1", u1)
		return err
	}
	return nil
}
func GetUsersByPagination() ([]User, error) {
	us := make([]User, 0)
	return us, nil
}

func PostUser(u User) error {
	if err := utils.DbUtils().Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func PutUser(u User) error {
	// update前需要使用Model(&User{}),选中需要更新的列
	if err := utils.DbUtils().Model(&User{ID: u.ID}).Update(&u).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(u User) error {
	fmt.Println("0000", u.ID)
	// 删除对象时需要指定主键,否则会触发批量delete
	//if err := utils.DbUtils().Unscoped().Delete(&u).Error; err != nil {
	if err := utils.DbUtils().Delete(&u).Error; err != nil {
		return err
	}
	return nil
}

func UserIsExisted(u User) (bool, error) {
	// 当err不是gorm.ErrRecordNotFound时,默认该用户已存在
	err := utils.DbUtils().Where(&u).First(&User{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return true, err
	}
	return true, nil
}

func UserNameIsExisted(u User) (bool, error, User) {
	// 当err不是gorm.ErrRecordNotFound时,默认该用户已存在
	fmt.Println(u.NickName)
	existedUser := User{}
	err := utils.DbUtils().Where(&User{NickName: u.NickName}).First(&existedUser).Error
	fmt.Println("-----", err)
	fmt.Println(existedUser)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil, existedUser
		}
		return true, err, existedUser
	}
	return true, nil, existedUser
}
