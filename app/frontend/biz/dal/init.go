package dal

import (
	"github.com/GreysonCarlos/gomall/app/frontend/biz/dal/mysql"
	"github.com/GreysonCarlos/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
