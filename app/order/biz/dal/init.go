package dal

import (
	"github.com/GreysonCarlos/gomall/app/order/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
