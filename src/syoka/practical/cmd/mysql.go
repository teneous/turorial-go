package cmd

import (
	"flag"
	"fmt"
	"os"
	"syoka/practical/utils"
)

func MysqlInput() {
	//退出程序
	_ = flag.NewFlagSet("mysql", flag.ExitOnError)
	flag.Usage = func() {
		helpInfo := `mysql cli tool.args:
		-h host mysql host
		-p port mysql prt 
		-P password mysql password`
		fmt.Println(helpInfo)
	}

	args := os.Args
	len := utils.ArrayStringLen(args)

	if len == 1 {
		flag.Usage()
		return
	}

	//var *host string()
	for index, _ := range args {
		switch os.Args[index+1] {
		case "h":
			flag.String("h", "127.0.0.1", "host of mysql")

			//database.Parse(os.Args[index+1])
		case "p":
			//port = flag.String("p", "3306", "host of mysql")
		case "P":
			//pwd = flag.String("P", "scout", "password of mysql")
			//fmt.Printf("%s -h %s -p %s -P %s", database, host, port, pwd)
		}
	}
}
