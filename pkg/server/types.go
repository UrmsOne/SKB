/*
@Author: urmsone urmsone@163.com
@Date: 2/23/21 8:47 PM
@Name: types.go
*/
package server

import "github.com/sirupsen/logrus"

type server struct {
	lg         logrus.FieldLogger
	cmdOptions CmdOptions
	stopCh     chan struct{}
}

type CmdOptions struct {
	Host string
	Post string
	Config
}

type Config struct {
	APP struct{}

	DataBase struct {
		Username, Password, Host, Port, DbName string
	}
}
