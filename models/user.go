package models

import "time"

type User struct {
	Id int
	Name string
	Age int
	Addr string
	Pic string
	Phone int
}

type UserInfo struct {
	Id int
	Name string
	CreateTime time.Time
}
