package models

import (
	"time"
)

type Shoppingcart struct {
	Createtime  time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	Goodsid     int       `xorm:"comment('商品id') INT"`
	Goodsnumber int       `xorm:"comment('商品数量') INT"`
	Id          int       `xorm:"not null pk autoincr comment('主键') INT"`
	IsDel       int       `xorm:"comment('0表示无效，1表示有效') INT"`
	Updatetime  time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Userid      int       `xorm:"INT"`
}
