package databases

import (
	"github.com/go-xorm/xorm"
)

var DBClient *xorm.Engine

func init()  {

	// DBClient = config.InitDB()
}

