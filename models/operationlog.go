package models

import (
	"time"
)

type Operationlog struct {
	Createtime    time.Time `xorm:"TIMESTAMP"`
	Id            int       `xorm:"not null pk autoincr comment('主键') INT"`
	IsDel         int       `xorm:"comment('0表示无效，1表示有效') INT"`
	Operationtype string    `xorm:"comment('操作类型') VARCHAR(255)"`
	Request       string    `xorm:"VARCHAR(255)"`
	Result        string    `xorm:"VARCHAR(255)"`
	Updatetime    time.Time `xorm:"TIMESTAMP"`
}
