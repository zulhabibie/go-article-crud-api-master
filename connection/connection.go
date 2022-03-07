package connection

import (
	"go-crud-article/structs"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
	Err error
)

func Connect()  {
	DB, Err = gorm.Open("mysql", "zulfikar:QEMcrrqfWItyQX9E@/zulfikar?charset=utf8&parseTime=True")
	
	if Err != nil {
		log.Println("Connection failed", Err)
		} else {
			log.Println("Server up and running")
		}
		
		DB.AutoMigrate(&structs.User{}, &structs.Risk_profile{})

}
