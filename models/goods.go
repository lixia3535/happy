package models

import (
	"time"
)

type Goods struct {
	Createtime   time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	Describe     string    `xorm:"comment('商品描述') TEXT"`
	Id           int       `xorm:"not null pk autoincr comment('主键') INT"`
	IsDel        int       `xorm:"comment('0表示无效，1表示有效') INT"`
	Name         string    `xorm:"comment('商品名字') VARCHAR(255)"`
	Price        float64   `xorm:"comment('商品价格') DOUBLE"`
	Productimage string    `xorm:"comment('商品图片') VARCHAR(255)"`
	Updatetime   time.Time `xorm:"comment('更新时间') TIMESTAMP"`
}
