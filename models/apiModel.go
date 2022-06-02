package models

import (
	"github.com/aWildProgrammer/fconf"
	"github.com/go-xorm/xorm"
)

// 将数据库连接实例，配置文件，gin.Context放入一个一个结构体里面

type ApiModel struct {
	Conf *fconf.Config
	DB *xorm.Engine
}

