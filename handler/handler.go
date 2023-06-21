// Package handler @Author:冯铁城 [17615007230@163.com] 2023-06-20 15:39:25
package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/thedevsaddam/gojsonq/v2"
)

// AddHandler 添加方法handler
func AddHandler(w http.ResponseWriter, req *http.Request) {

	//1.从body获取参数
	if data, err := io.ReadAll(req.Body); err == nil {
		url := gojsonq.New().FromString(string(data)).Find("url").(string)

		//2.定义key
		var key string

		//3.存入map
		if err := store.Set(&url, &key); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			fmt.Fprintf(w, "%v:%v/%v", *host, *port, key)
		}
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// RedirectHandler 重定向handler
func RedirectHandler(w http.ResponseWriter, req *http.Request) {

	//1.获取key
	key := req.URL.Path[1:]

	//2.定义url
	var url string

	//3.获取url
	if err := store.Get(&key, &url); err == nil {
		http.Redirect(w, req, url, http.StatusFound)
	} else {
		http.NotFound(w, req)
	}
}

// InitHandler 初始化handler
func InitHandler(f func(w http.ResponseWriter, req *http.Request)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		//1.defer处理异常
		defer func() {
			if err := recover(); err != nil {
				log.Printf("handler running error->[%v]", err)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		//2.执行参数方法
		f(writer, request)
	}
}
