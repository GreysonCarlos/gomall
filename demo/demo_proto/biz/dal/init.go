package dal

import (
	"github.com/GreysonCarlos/demo/demo_proto/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
