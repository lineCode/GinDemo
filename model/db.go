package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Gorm is goroutines friendly, so you can create a global variable to keep the connection and use it everywhere like this

var DB gorm.DB
var err error

func init() {
	DB, err = gorm.Open("mysql", "root:1234@/gin_demo?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("connect database success")
	}

	DB.DB()
	DB.DB().Ping()
	DB.DB().SetMaxIdleConns(50)  //闲（等待）
	DB.DB().SetMaxOpenConns(200) //打开
	DB.SingularTable(true)       //单数表

	// DB.AutoMigrate(&User{},&Article{},&Comment{})  //自动迁移

	// DB.Model(&Article{}).AddForeignKey("user_id", "user(id)", "RESTRICT", "RESTRICT")   //RESTRICT，限制
	// DB.Model(&Comment{}).AddForeignKey("user_id", "user(id)", "RESTRICT", "RESTRICT")
	// DB.Model(&Comment{}).AddForeignKey("article_id", "article(id)", "RESTRICT", "RESTRICT")
}
