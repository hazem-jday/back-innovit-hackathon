// @/entities/publication.go
package entities

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Id              string `json:"id"`
	Title           string `json:"titre"`
	Image           string `json:"image"`
	DatePublication string `json:"datePublication"`
}
