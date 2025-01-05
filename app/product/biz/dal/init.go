package dal

import (
	"github.com/GreysonCarlos/gomall/app/product/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
