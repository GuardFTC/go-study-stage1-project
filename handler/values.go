// Package handler @Author:冯铁城 [17615007230@163.com] 2023-06-20 17:17:42
package handler

import (
	"flag"

	"go-study-stage1-project/structs"
)

// 定义us变量
var store structs.Store

// 定义部分常量参数
var (
	host       = flag.String("host", "localhost", "the server start host")
	port       = flag.String("port", "8080", "the server start port")
	rpcEnabled = flag.Bool("rpc", false, "is server support rpc")
	masterAddr = flag.String("master", "", "the master server address")
)
