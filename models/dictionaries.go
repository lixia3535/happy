package models

import (
	"time"
)

/**
dictionnaries 数据字典
*/
type Dictionaries struct {
	Createtime time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	Haschild   int       `xorm:"comment('是否有子节点 0表示没有，1表示有') INT"`
	Id         int       `xorm:"not null pk autoincr comment('主键') INT"`
	IsDel      int       `xorm:"comment('0表示无效，1表示有效') INT"`
	Name       string    `xorm:"comment('名称') VARCHAR(255)"`
	Parentid   int       `xorm:"comment('parentId') INT"`
	Property   string    `xorm:"VARCHAR(255)"`
	Seq        int       `xorm:"comment('序号') INT"`
	Updatetime time.Time `xorm:"comment('更新时间') TIMESTAMP"`
}
