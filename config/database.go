// @/config/database.go
package config

import (
	"back/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "root:root@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
	var err error

	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.User{})
	Database.AutoMigrate(&entities.Post{})

	return nil
}
