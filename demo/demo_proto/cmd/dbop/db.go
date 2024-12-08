package main

import (
	// "fmt"

	"github.com/GreysonCarlos/demo/demo_proto/biz/dal"
	"github.com/GreysonCarlos/demo/demo_proto/biz/dal/mysql"
	"github.com/GreysonCarlos/demo/demo_proto/model"
	"github.com/joho/godotenv"
)
func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	// CURD

	// create
	// mysql.DB.Create(&model.User{Email: "demo@example", Password: "demo1"})
	
	// update
	// mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example").Update("password", "demo2")

	// Read
	// var row model.User
	// mysql.DB.Model(&model.User{}).Where("email = ?", "demo@example").First(&row)

	// fmt.Printf("row: %+v\n", row)

	// delete
	// mysql.DB.Where("email = ?", "demo@example").Delete(&model.User{})
	mysql.DB.Unscoped().Where("email = ?", "demo@example").Delete(&model.User{})
}