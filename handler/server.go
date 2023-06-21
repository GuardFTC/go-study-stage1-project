// Package handler @Author:冯铁城 [17615007230@163.com] 2023-06-20 17:18:11
package handler

import (
	"log"
	"net/http"
	"net/rpc"

	"go-study-stage1-project/structs"
)

// StartServer 开启服务器
func StartServer() {

	//1.初始化store
	if *masterAddr == "" {
		store = structs.InitUrlStore()
	} else {
		store = structs.InitClientStore(*masterAddr)
	}

	//2.判定是否开启rpc服务
	if *rpcEnabled {
		rpc.RegisterName("Store", store)
		rpc.HandleHTTP()
	}

	//3.加载handler
	http.HandleFunc("/add", InitHandler(AddHandler))
	http.HandleFunc("/", InitHandler(RedirectHandler))

	//4.开启服务器
	if err := http.ListenAndServe(*host+":"+*port, nil); err != nil {
		log.Panicf("server error->[%v]", err)
	}
}
