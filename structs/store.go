// Package structs @Author:冯铁城 [17615007230@163.com] 2023-06-21 09:47:23
package structs

type Store interface {
	Set(url, key *string) error
	Get(key, url *string) error
}
