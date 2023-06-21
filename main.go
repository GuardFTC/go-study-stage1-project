// @Author:冯铁城 [17615007230@163.com] 2023-06-20 16:09:35
package main

import (
	"flag"

	"go-study-stage1-project/handler"
)

func main() {

	//1.加载flag
	flag.Parse()

	//2.开启服务器
	handler.StartServer()
}
