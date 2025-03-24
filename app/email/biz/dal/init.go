package dal

import (
	"github.com/GreysonCarlos/gomall/app/email/biz/dal/mysql"
	"github.com/GreysonCarlos/gomall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
