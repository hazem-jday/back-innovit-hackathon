// @/entities/user.go
package entities

// Modèle des utilisateurs
type User struct {
	Username      string `json:"username"`
	Id            int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Nom           string `json:"nom"`
	Prenom        string `json:"prenom"`
	Email         string `json:"email"`
	Sexe          string `json:"sexe"`
	ImageProfil   string `json:"imageProfil"`
	Password      string `json:"password"`
	DateNaissance string `json:"dateNaissance"`
	Telephone     string `json:"telephone"`
}

// Modèle de Login utilisateur
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
