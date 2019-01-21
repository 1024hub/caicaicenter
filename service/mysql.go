package service

import (
	"fmt"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"log"
	"time"
)

// engine package db
var engine *xorm.Engine

func 	Initdb()  {
 	var err error
	//创建orm引擎
	cfg, err := ini.Load("src/yibiancheng.com/accountcenter/config/db.ini")
		if err != nil {
			fmt.Println(err)
		}
		username := cfg.Section("mysql").Key("username").Value()
		password := cfg.Section("mysql").Key("password").Value()
		url := cfg.Section("mysql").Key("url").Value()

		source := fmt.Sprintf("%s:%s@%s", username, password, url)
		engine, err = xorm.NewEngine("mysql", source)
	//engine, err = xorm.NewEngine("mysql", "root:12345678@/test?charset=utf8")
	defer engine.Close()

	if err!=nil{
		fmt.Println(err)
		return
	}
	//连接测试
	if err := engine.Ping(); err!=nil{
		fmt.Println(err)
		return
	}
	//	/*--------------------------------------------------------------------------------------------------
	//	1、使用RegisterSqlMap()注册SqlMap配置
	//	2、RegisterSqlTemplate()方法注册SSqlTemplate模板配置
	//	3、SqlMap配置文件总根目录和SqlTemplate模板配置文件总根目录可为同一目录
	//	--------------------------------------------------------------------------------------------------*/
	//
	//	//注册SqlMap配置，可选功能，如应用中无需使用SqlMap，可无需初始化
	//	//此处使用xml格式的配置，配置文件根目录为"./sql/oracle"，配置文件后缀为".xml"
	//	err = x.RegisterSqlMap(xorm.Xml("./db/sql/xml", ".xml"))
	//	if err != nil {
	//		log.Fatalf("db error: %#v\n", err.Error())
	//	}
	//	//注册动态SQL模板配置，可选功能，如应用中无需使用SqlTemplate，可无需初始化
	//	//此处注册动态SQL模板配置，使用Pongo2模板引擎，配置文件根目录为"./sql/oracle"，配置文件后缀为".stpl"
	//	err = x.RegisterSqlTemplate(xorm.Default("./db/sql/tpl", ".sql"))
	//	if err != nil {
	//		log.Fatalf("db error: %#v\n", err.Error())
	//	}
	//	//开启SqlMap配置文件和SqlTemplate配置文件更新监控功能，将配置文件更新内容实时更新到内存，如无需要可以不调用该方法
	//	err = x.StartFSWatcher()
	//	if err != nil {
	//		log.Printf("sql parse error: %#v\n", err)
	//	}
	//日志打印SQL
	engine.ShowSQL(true)

	//设置连接池的空闲数大小
	engine.SetMaxIdleConns(5)
	//设置最大打开连接数
	engine.SetMaxOpenConns(5)
		// 30minute ping db to keep connection
		timer := time.NewTicker(time.Minute * 30)
		go func(x *xorm.Engine) {
			for _ = range timer.C {
				err = x.Ping()
				if err != nil {
					log.Fatalf("db connect error: %#v\n", err.Error())
				}
			}
		}(engine)
		//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
		//engine.SetTableMapper(core.SnakeMapper{})

		//增
		//user:=&entity.User{
		//	Username:"zhangsan",
		//}
		//affected,err := engine.Insert(user)
		//fmt.Println(affected)
		//
		//删
		//user:=&entity.User{
		//	Username:"zhangsan",
		//}
		//affected_delete,err := engine.Delete(user)
		//fmt.Println(affected_delete)

		//改
		//user := new(User)
		//user.Username="tyming"
		//affected_update,err := engine.Id(1).Update(user)
		//fmt.Println(affected_update)

		//查
		//user := new(User)
		//\result,err := engine.Id(1).Get(user)
		//result,err := engine.Where("id=?",1).Get(user)
		//fmt.Println(result)
}