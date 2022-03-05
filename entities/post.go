// @/entities/publication.go
package entities

//Modèle du Post
type Post struct {
	Id              int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title           string `json:"titre"`
	Description     string `json:"description"`
	Username        string `json:"username"`
	Location        string `json:"location"`
	Image           string `json:"image"`
	DatePublication string `json:"datePublication"`
}
