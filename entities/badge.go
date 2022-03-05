// @/entities/badge.go
package entities

// Modèle de Badge
type Badge struct {
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Type string `json:"type"`
	Name string `json:"name"`
}
