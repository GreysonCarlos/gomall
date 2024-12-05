package dal

import (
	"github.com/GreysonCarlos/gomall/biz/dal/mysql"
	"github.com/GreysonCarlos/gomall/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
