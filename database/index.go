package database

import (
	"gorm.io/driver/postgres"
  "gorm.io/gorm"
	"github.com/paafff/golang-CRUD/models"
)

var DB *gorm.DB

func ConnectDatabase()  {
dsn := "host=localhost user=postgres password=paafff dbname=dbgolang_crud port=5432 sslmode=disable TimeZone=Asia/Shanghai"
//  dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

	database,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
// return err
	}

	database.AutoMigrate(&models.Product{})
	
DB = database

}