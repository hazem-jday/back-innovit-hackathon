// @/entities/notification.go
package entities

// Modèle de Notification
type Notification struct {
	Id          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string `json:"titre"`
	Description string `json:"description"`
	Date        string `json:"date"`
}
