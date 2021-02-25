## echo sever-demo

### 启动

项目架构
```
.
├── cmd
│   ├── main.go
│   └── runServer.go
├── conf
│   ├── air.conf
│   └── config.yaml
├── deploy
├── Dockerfile
├── go.mod
├── go.sum
├── log
│   └── server.log
├── Makefile
├── pkg
│   ├── handler
│   │   └── user
│   │       └── handler.go
│   ├── middleware
│   │   └── echoLogrus
│   │       └── middleware.go
│   ├── model
│   │   ├── types.go
│   │   └── user.go
│   ├── server
│   │   ├── server.go
│   │   └── types.go
│   ├── service
│   │   └── user
│   │       ├── service.go
│   │       └── types.go
│   └── utils
│       ├── config.go
│       ├── db.go
│       ├── logger.go
│       ├── mystrconv
│       │   └── mystrconv.go
│       └── pprof
│           └── pprof.go
├── readme.md
├── server
└── test
    └── db
        └── gorm.go
```

* cmd: 项目入口文件
* conf: 项目配置文件
    * config.yaml: 项目启动的相关配置,如数据库连接信息等
    * air.conf: echo开发环境中热重启的配置信息
* deploy: k8s部署文件(kustomize风格)
* log: 默认日志输出目录
* pkg: 业务逻辑处理
    * handler: restful服务路由控制器(类似于java的controller层),根据url调用service层对应的路由函数处理业务逻辑
    * middleware: 存放echo server各种中间件,包括日志中间件,权限验证中间件等(类似于flask的装饰器或各种钩子函数)
    * model: ORM对应的结构体定义,以及对数据库的各种查询操作(类似于java的Dao层),被service层调用
    * server: echo server的入口,被cmd目录的runServer.go调用
    * service: 编写各种业务逻辑,依赖model层(类似于java的service层),被handler层调用
    * utils: 封装项目通用的工具,如日志初始化,db初始化,k8s client初始化,redis初始化等
* test: 存放项目测试代码
* Dockerfile: 打包docker镜像
* Makefile: 项目部署
* readme.md

项目完整调用链路

`browser->handler->service->model->db(mysql,redis,mongodb)`

开发流程

* 首先在handler中创建相关服务的handler文件,并设置相关路由, 路由控制函数指向同名service中的路由函数;
```
/*
@Author: urmsone urmsone@163.com
@Date: 2/24/21 1:37 PM
@Name: handler.go
*/
package user

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"SKB/pkg/service/user"
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
```
* 在service中创建对应的service接口,并编写相关实现

```
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
```

* 在model中创建ORM结构体,编写数据库相关操作函数

```
type User struct {
	// gorm.Model
	ID        uint `gorm:"primary_key"`
    ...
}

func GetUserByID(id string) (*User, error) {
	u := &User{}
	if err := utils.DbUtils().First(&u, id).Error; err != nil {
		return nil, err
	}
	return u, nil
}
```

加载配置文件

* viper+cobra
* 命令行参数　>　环境变量　>　配置文件

swagger集成
* TODO

日志
* logrus与echo集成: 详情请看pkg/middleware/echoLogrus
* logrus与lumberjack集成实现滚动日志: 详情请看pkg/utils/logger.go

ORM
* GORM集成

集成开发环境热重启
* 使用air工具

运行

* 开发调试

开发环境:
1) Ubuntu 18.04
2) go 1.14+
3) mysql (docker image mysql:8)
```
1) go mod download
2) go get github.com/cosmtrek/air
3) 启动数据库, `docker run -itd --name SKB -e MYSQL_ROOT_PASSWORD="123456" -p 3310:3306 mysql:8 --datadir /var/lib/mysql/datadir`
4) 创建数据库表,配置数据库连接信息
5) air -c conf/air.conf(air负责监控文件的变化,实现热重启,win下不一定可行)
root@ubuntu:/workspace/gospace/SKB# air -c conf/air.conf 

  __    _   ___  
 / /\  | | | |_) 
/_/--\ |_| |_| \_ v1.12.1 // live reload for Go apps, with Go1.14.0

mkdir /workspace/gospace/SKB/tmp
watching .
watching cmd
watching conf
watching deploy
!exclude log
watching pkg
watching pkg/handler
watching pkg/handler/user
watching pkg/middleware
watching pkg/middleware/echoLogrus
watching pkg/model
watching pkg/server
watching pkg/service
watching pkg/service/user
watching pkg/utils
watching pkg/utils/mystrconv
watching pkg/utils/pprof
watching test
watching test/db
!exclude tmp
building...
running...
timestamp="2021-02-25T18:21:37Z" level=info msg="config.database <nil>"
timestamp="2021-02-25T18:21:37Z" level=info msg="[database.password database.dbname database.host database.port database.username]"
timestamp="2021-02-25T18:21:37Z" level=info msg="Server starting ..."

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.2.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1234
```

### 正式环境部署

TODO

### 相关文档
echo[官网文档](https://echo.labstack.com/guide)

gorm[中文文档](https://gorm.io/zh_CN/docs/index.html)

air [github链接](https://github.com/cosmtrek/air)