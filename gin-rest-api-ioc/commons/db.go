package  commons;

import (
	"fmt"
	"gin-rest-api-ioc/model"
	"github.com/jinzhu/gorm"
)

import _ "github.com/jinzhu/gorm/dialects/postgres"



var db *gorm.DB;



func init()  {
	//db1, err := gorm.Open("sqlite3", "d:/test.db")


	db1, err := gorm.Open("postgres", "host=192.168.1.101 port=5432 user=postgres dbname=postgres password=shasha123456 sslmode=disable")
	if err != nil {
		fmt.Println(err);
		panic("failed to connect database")
	}
	db=db1;
	db.AutoMigrate(&model.User{})

	db.Create(&model.User{UserName: "qingtian"})
	db.Create(&model.User{UserName: "zzsoszz"})

	fmt.Println("db inited");
}


func GetDb()  *gorm.DB {
	return db;
}

