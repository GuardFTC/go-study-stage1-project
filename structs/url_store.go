// Package structs @Author:冯铁城 [17615007230@163.com] 2023-06-20 14:55:16
package structs

import (
	"errors"
	"log"
	"sync"

	"go-study-stage1-project/utils"
)

// UrlStore Url存储结构体
type UrlStore struct {
	urlMap map[string]string
	mu     sync.RWMutex
}

// InitUrlStore 初始化Url存储
func InitUrlStore() *UrlStore {
	return &UrlStore{urlMap: make(map[string]string)}
}

// Set 设置键值对方法
func (us *UrlStore) Set(url, key *string) error {
	log.Println("master server set")

	//1.加锁
	us.mu.Lock()
	defer us.mu.Unlock()

	//2.生成随机key
	if *key == "" {
		*key = utils.RandomString(7)
	}

	//3.存值
	if _, isPresent := us.urlMap[*key]; !isPresent {
		us.urlMap[*key] = *url
	}

	//4.返回
	return nil
}

// Get 取值方法
func (us *UrlStore) Get(key, url *string) error {
	log.Println("master server get")

	//1.加锁
	us.mu.RLock()
	defer us.mu.RUnlock()

	//2.取值
	if value, isPresent := us.urlMap[*key]; isPresent {
		*url = value
		return nil
	} else {
		return errors.New("can't find url by this key")
	}
}
