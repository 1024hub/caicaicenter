package entity

import "time"


//
////数据库表明是demo_user 但是结构体里对应的是User表
//func (u User) TableName() string {
//	return "demo_user"
//}
type User struct {
	Id       int       `xorm:"not null pk autoincr INT(11)"`
	Username string    `xorm:"not null VARCHAR(32)"`
	Birthday time.Time `xorm:"DATE"`
	Sex      string    `xorm:"CHAR(1)"`
	Address  string    `xorm:"VARCHAR(256)"`
}