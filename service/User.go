package service

import (
	"fmt"
	"yibiancheng.com/accountcenter/entity"
)

func SaveUsers( p *entity.User) {
 	affected,err := engine.Insert(p)
	if err!=nil{
		fmt.Println(affected)
	}
}