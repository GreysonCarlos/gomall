package dal

import (
	"github.com/GreysonCarlos/projects/Gomall/app/frontend/biz/dal/mysql"
	"github.com/GreysonCarlos/projects/Gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
