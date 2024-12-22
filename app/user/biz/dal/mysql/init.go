package mysql

import (
	"fmt"
	"os"

	"github.com/GreysonCarlos/gomall/app/user/model"
	// "github.com/GreysonCarlos/gomall/app/user/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3307)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),	
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&model.User{})

	if err != nil {
		panic(err)
	}

	type Version struct {
		Version	string
	}

	var v Version
	err = DB.Raw("select version() as version").Scan(&v).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(v)
}
