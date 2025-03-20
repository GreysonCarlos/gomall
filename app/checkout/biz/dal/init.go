package dal

import (
	"github.com/GreysonCarlos/gomall/app/checkout/biz/dal/mysql"
	"github.com/GreysonCarlos/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
