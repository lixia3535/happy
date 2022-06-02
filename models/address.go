package models

import (
	"time"
)

type Address struct {
	Cityid         int       `xorm:"comment('市') INT"`
	Createtime     time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	Detailedddress string    `xorm:"comment('详细地址') VARCHAR(255)"`
	Id             int       `xorm:"not null pk autoincr comment('主键') INT"`
	IsDel          int       `xorm:"comment('0表示无效，1表示有效') INT"`
	Provinceid     int       `xorm:"comment('省份') INT"`
	Tel            string    `xorm:"comment('电话') VARCHAR(15)"`
	Updatetime     time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Userid         int       `xorm:"comment('userId') INT"`
}
