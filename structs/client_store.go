// Package structs @Author:冯铁城 [17615007230@163.com] 2023-06-21 09:26:34
package structs

import (
	"log"
	"net/rpc"
)

// ClientStore 客户端Store结构体
type ClientStore struct {
	client *rpc.Client
	us     *UrlStore
}

// InitClientStore 初始化客户端Store
func InitClientStore(addr string) *ClientStore {

	//1.初始化结构体
	clientStore := new(ClientStore)

	//2.设置urlStore
	clientStore.us = InitUrlStore()

	//3.连接主服务器
	if client, err := rpc.DialHTTP("tcp", addr); err != nil {
		log.Printf("init client_store rpc error->[%v]", err)
	} else {
		clientStore.client = client
	}

	//4.返回
	return clientStore
}

// Set 设置键值对方法
func (c *ClientStore) Set(url, key *string) error {
	log.Println("client server set")
	return c.client.Call("Store.Set", url, key)
}

// Get 取值方法
func (c *ClientStore) Get(key, url *string) error {
	log.Println("client server get")

	//1.从本地map取值
	if err := c.us.Get(key, url); err == nil {
		return nil
	}

	//2.本地未获取到，调用主服务器取值
	if err := c.client.Call("Store.Get", key, url); err != nil {
		return err
	} else {
		c.us.Set(url, key)
		return nil
	}
}
