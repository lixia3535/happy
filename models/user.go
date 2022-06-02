package models

import (
	"time"
)

type User struct {
	Createtime time.Time `xorm:"comment('创建时间') TIMESTAMP" json:"createtime"`
	Describe   string    `xorm:"comment('简介') VARCHAR(255)" json:"describe"`
	Icon       string    `xorm:"comment('头像') VARCHAR(255)" json:"icon"`
	Id         int       `xorm:"not null pk autoincr comment('主键') INT" json:"id"`
	IsDel      int       `xorm:"comment('0表示无效，1表示有效') INT" json:"is_del"`
	Password   string    `xorm:"comment('密码') VARCHAR(255)" json:"password"`
	Updatetime time.Time `xorm:"comment('修改时间') TIMESTAMP" json:"updatetime"`
	Username   string    `xorm:"comment('用户名') VARCHAR(100)" json:"username"`
}

