package dal

import (
	"github.com/GreysonCarlos/gomall/app/cart/biz/dal/mysql"
	"github.com/GreysonCarlos/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
