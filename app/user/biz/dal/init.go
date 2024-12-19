package dal

import (
	"github.com/GreysonCarlos/gomall/app/user/biz/dal/mysql"
	"github.com/GreysonCarlos/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
