package models

type Proj struct {
	// Model
	ID    uint   `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Users []User `gorm:"many2many:user_proj;" json:"users"`
}
