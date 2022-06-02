package models

import (
	"time"
)

type Orderform struct {
	Addressid   int       `xorm:"comment('地址') INT"`
	Createtime  time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	Goodsid     int       `xorm:"comment('goodsId') INT"`
	Goodsnumber int       `xorm:"comment('商品数量') INT"`
	Goodsprice  float64   `xorm:"comment('商品价格') DOUBLE"`
	Id          int       `xorm:"not null pk autoincr comment('主键') INT"`
	IsDel       int       `xorm:"comment('0表示无效，1表示有效') INT"`
	Ordernumber string    `xorm:"comment('订单编号') VARCHAR(255)"`
	Status      int       `xorm:"comment('订单状态 0表示待发货，1表示已经发货，3表示已收货，4表示退货') INT"`
	Updatetime  time.Time `xorm:"comment('更新时间') TIMESTAMP"`
	Userid      int       `xorm:"comment('userid') INT"`
}
