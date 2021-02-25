/*
@Author: urmsone urmsone@163.com
@Date: 2/23/21 8:26 PM
@Name: runServer.go
*/
package main

import (
	"SKB/pkg/server"
	"SKB/pkg/utils"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

var runServerCfg = struct {
	serverHost string
	serverPort string
	logDir     string
	configDir  string
}{}

func init() {

	runServerCmd.Flags().StringVar(&runServerCfg.serverHost, "server.host", "", "")
	runServerCmd.Flags().StringVar(&runServerCfg.serverPort, "server.port", "1234", "")
	runServerCmd.Flags().StringVar(&runServerCfg.logDir, "server.logDir", "log", "")
	runServerCmd.Flags().StringVar(&runServerCfg.configDir, "config.dir", "conf", "")

	//cobra.OnInitialize()

	rootCmd.AddCommand(runServerCmd)
}

var runServerCmd = &cobra.Command{
	Use:   "run",
	Short: "run server",
	Long:  "Run server, listening host:port and providing service",
	RunE:  runServe,
	// Hook before and after Run initialize and do something, respectively
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	//BashCompletionFunction: ``
}

func runServe(cmd *cobra.Command, args []string) error {
	fmt.Println(`
`)
	if err := cmd.Flags().Parse(args); err != nil {
		return err
	}
	// 初始化logger
	lg := utils.NewLog(runServerCfg.logDir)
	if lg == nil {
		// TODO: 实现自定义error
		return nil
	}
	// 加载配置文件
	var config server.Config
	if err := utils.ConfigInit("", "", config); err != nil {
		lg.Errorln(err)
		return err
	}
	// TODO: 验证配置文件字段,验证命令行参数
	lg.Println("config.database", utils.GetConfig(&config))
	lg.Println(utils.ConfigUtils().AllKeys())

	utils.ConfigUtils().OnConfigChange(func(in fsnotify.Event) {
		// TODO: 配置文件热加载,服务优雅重启
		lg.Println("configFile changed ", in.Name, in.Op)
	})
	utils.ConfigUtils().WatchConfig()
	stopCh := make(chan struct{})
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	g := errgroup.Group{}
	g.Go(func() error {
		// run server
		lg.Infof("Server starting ...")
		ops := server.CmdOptions{
			Host:   runServerCfg.serverHost,
			Post:   runServerCfg.serverPort,
			Config: config,
		}
		apiServer := server.NewServer(lg, ops, stopCh)
		//return server.Run(lg, ops, stopCh)
		return apiServer.Run(lg, ops, stopCh)
	})
	g.Go(func() error {
		// 监听信号,控制server的退出
		for {
			select {
			//case s := <- signalCh:
			case s := <-signalCh:
				lg.Infof("Captured %s signal. Exiting ...\n", s)
				close(stopCh)
				lg.Infof("Server stopping ...")
				return nil
			}
		}
	})

	return g.Wait()
}
