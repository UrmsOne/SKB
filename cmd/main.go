/*
@Author: urmsone urmsone@163.com
@Date: 2/23/21 8:20 PM
@Name: main.go
*/
package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "Server demo",
	Short: "Server example",
	Long: `Server Endpoint xxx
and etc.`,
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
