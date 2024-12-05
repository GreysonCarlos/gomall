package dal

import (
	"github.com/GreysonCarlos/gomall/demo/demo_thrift/biz/dal/mysql"
	"github.com/GreysonCarlos/gomall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
