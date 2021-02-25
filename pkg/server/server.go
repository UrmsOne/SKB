/*
@Author: urmsone urmsone@163.com
@Date: 2/23/21 8:47 PM
@Name: server.go
*/
package server

import (
	"SKB/pkg/handler/user"
	"SKB/pkg/middleware/echoLogrus"
	"SKB/pkg/model"
	"SKB/pkg/utils"
	"SKB/pkg/utils/pprof"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/echo-swagger"
	"net/http"
	"time"
)

func (s *server) Run(lg logrus.FieldLogger, ops CmdOptions, stopCh chan struct{}) error {
	// 初始化db
	if err := utils.DbInit(ops.DataBase.Username, ops.DataBase.Password, ops.DataBase.Host, ops.DataBase.Port, ops.DataBase.DbName); err != nil {
		lg.Errorln("db init error: ", err)
		return err
	}

	utils.DbUtils().Debug().AutoMigrate(
		model.User{},
	)

	e := echo.New()
	e.Debug = true
	logger := echoLogrus.NewLoggerMiddleware(lg)
	e.Logger = logger
	e.Use(logger.Hook())
	e.Use(middleware.Recover())
	pprof.Wrap(e)
	hdl := e.Group("/api/v1")
	hdl.GET("/swagger/*", echoSwagger.EchoWrapHandler())
	user.RegisterHandler(hdl, lg)
	go func() {
		select {
		case <-stopCh:
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			s.lg.Infoln("Shutdown server error: ", e.Shutdown(ctx))
			lg.Infoln("Closing db connection ...")
			if err := utils.DbClose(); err != nil {
				lg.Errorln(err)
			}
			lg.Infoln("Db connection closed...")
			return
		}
	}()
	if err := e.Start(fmt.Sprintf("%s:%s", ops.Host, ops.Post)); err != nil {
		if err == http.ErrServerClosed {
			// 主动关闭server,如接收到控制台退出信号等...
			s.lg.Println("Http server closed, bye!")
			return nil
		} else {
			// 非主动关闭
			return err
		}
	}
	//return e.Start(fmt.Sprintf("%s:%s", ops.Host, ops.Post))
	//<-stopCh
	return nil
}

func NewServer(lg logrus.FieldLogger, ops CmdOptions, stopCh chan struct{}) *server {
	return &server{
		lg:         lg,
		cmdOptions: ops,
		stopCh:     stopCh,
	}
}
