/*
@Author: urmsone urmsone@163.com
@Date: 2/24/21 1:37 PM
@Name: handler.go
*/
package user

import (
	"SKB/pkg/service/user"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func RegisterHandler(e *echo.Group, lg logrus.FieldLogger) {
	svc := user.NewUserServiceImpl(lg)
	r := e.Group("/user")
	r.GET("", svc.GetList)
	r.GET("/:id", svc.Get)
	r.POST("", svc.Post)
	r.PUT("", svc.Put)
	r.DELETE("/:id", svc.Delete)
}
