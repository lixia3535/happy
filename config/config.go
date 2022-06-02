package config

import (
	"fmt"
	"github.com/aWildProgrammer/fconf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"xorm.io/core"
)

// 设置一个全局变量、
//

func InitConf() *fconf.Config{
	conf ,err := fconf.NewFileConf("./resource/conf.ini")
	if err != nil{
		log.Fatal("加载配置文件失败",err.Error())
		return nil
	}
	return conf
}

func InitDB() *xorm.Engine{
	var  DB *xorm.Engine
	var connErr error
	conf := InitConf()
	if conf == nil {
		log.Fatal("加载配置文件失败")
		return nil
	}
	// 获取配置文件中的参数
	dbtype := conf.String("mysql.db1.type")
	databaseName := conf.String("mysql.db1.databaseName")
	host := conf.String("mysql.db1.Host")
	port := conf.String("mysql.db1.Port")
	pwd := conf.String("mysql.db1.Pwd")
	username := conf.String("mysql.db1.Username")

	//拼接参数
	DB, connErr = xorm.NewEngine(dbtype, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",username,pwd,host,port,databaseName))

	if connErr != nil{
		log.Fatal(connErr.Error())
	}

	//测试查看是否能联通数据库
	err := DB.Ping()
	if err != nil {
		log.Fatal("数据库连接失败！")
		return nil
	}
	// 显示打印sql语句
	DB.ShowSQL(true)
	// 设置连接池的空闲数大小
	DB.SetMaxIdleConns(5)
	// 设置打开的最大连接数
	DB.SetMaxOpenConns(5)

	//名称映射规则主要负责结构体名称到表名和结构体feild到表字段的名称映射
	DB.SetTableMapper(core.SnakeMapper{})

	return DB
}


