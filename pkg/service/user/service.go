/*
@Author: urmsone urmsone@163.com
@Date: 2/24/21 1:46 PM
@Name: service.go
*/
package user

import (
	"SKB/pkg/model"
	"SKB/pkg/utils/mystrconv"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type userService interface {
	Get(ctx echo.Context) error
	GetList(ctx echo.Context) error
	Post(ctx echo.Context) error
	Put(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type UserServiceImpl struct {
	lg logrus.FieldLogger
}

func (s *UserServiceImpl) Get(ctx echo.Context) error {
	id := ctx.Param("id")
	user, err := model.GetUserByID(id)
	if err != nil {
		return ctx.JSON(http.StatusOK, echo.Map{"err": err, "msg": "find user failed"})
	}
	// do something
	s.lg.Infoln(user)
	return ctx.JSON(http.StatusOK, user)
}

func (s *UserServiceImpl) GetList(ctx echo.Context) error {
	if users, err := model.GetUsers(); err != nil {
		return ctx.JSON(http.StatusOK, map[string]string{"msg": err.Error()})
	} else {
		return ctx.JSON(http.StatusOK, users)
	}
}

func (s *UserServiceImpl) Post(ctx echo.Context) error {
	data := struct {
		Username string
		NickName string
	}{}
	if err := ctx.Bind(&data); err != nil {
		s.lg.Println(err)
	}
	s.lg.Println(data)
	u := model.User{UserName: data.Username, NickName: data.NickName}
	if ok, err, _ := model.UserNameIsExisted(u); ok {
		return ctx.JSON(http.StatusOK, echo.Map{"err": err, "msg": "Nickname is existed"})
	}
	if err := model.PostUser(u); err != nil {
		return ctx.JSON(http.StatusOK, echo.Map{"err": err, "msg": "Create user failed"})
	}
	return ctx.JSON(http.StatusCreated, echo.Map{"msg": "Create user success"})
}

func (s *UserServiceImpl) Put(ctx echo.Context) error {
	updateUser := model.User{}
	if err := ctx.Bind(&updateUser); err != nil {
		return ctx.JSON(http.StatusOK, echo.Map{"err": err, "msg": "binding request data failed"})
	}
	if ok, err, exUser := model.UserNameIsExisted(updateUser); ok && exUser.ID != updateUser.ID {
		return ctx.JSON(http.StatusOK, echo.Map{"err": err, "msg": "Nickname is existed"})
	}
	if err := model.PutUser(updateUser); err != nil {
		return ctx.JSON(http.StatusOK, echo.Map{"err": err, "msg": "put user failed"})
	}
	return ctx.JSON(http.StatusAccepted, echo.Map{"msg": "put user success"})
}

func (s *UserServiceImpl) Delete(ctx echo.Context) error {
	// 路径参数使用ctx.Param("keyname")取值
	id, err := mystrconv.ParseUint(ctx.Param("id"))
	fmt.Println("id: ", id)
	if err != nil {
		return ctx.JSON(http.StatusAccepted, echo.Map{"err": "id parse to uint failed", "msg": "delete user failed"})
	}

	if id == 0 {
		return ctx.JSON(http.StatusAccepted, echo.Map{"err": "id must be required", "msg": "delete user failed"})
	}
	if err := model.DeleteUser(model.User{ID: id}); err != nil {
		return ctx.JSON(http.StatusAccepted, echo.Map{"err": err, "msg": "delete user failed"})
	}
	//return ctx.NoContent(http.StatusNoContent)
	return ctx.JSON(http.StatusAccepted, echo.Map{"msg": "delete user success"})
}

func NewUserServiceImpl(lg logrus.FieldLogger) *UserServiceImpl {
	return &UserServiceImpl{lg: lg}
}
