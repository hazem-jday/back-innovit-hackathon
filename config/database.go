// @/config/database.go
package config

import (
	"back/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Go ORM Database
var Database *gorm.DB
var DATABASE_URI string = "user:pass@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"

//Connexion à la base
func Connect() error {
	var err error

	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	//Initialisation des tableaux
	Database.AutoMigrate(&entities.User{})
	Database.AutoMigrate(&entities.Post{})

	return nil
}
