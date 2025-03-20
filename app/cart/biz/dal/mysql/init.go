package mysql

import (
	"fmt"
	"os"

	"github.com/GreysonCarlos/gomall/app/cart/conf"
	"github.com/GreysonCarlos/gomall/app/cart/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	_ = DB.AutoMigrate(&model.Cart{})
	if err != nil {
		panic(err)
	}
}
