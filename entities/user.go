// @/entities/user.go
package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id            string `json:"id"`
	Nom           string `json:"nom"`
	Prenom        string `json:"prenom"`
	Login         string `json:"login"`
	Password      string `json:"password"`
	DateNaissance string `json:"dateNaissance"`
}
