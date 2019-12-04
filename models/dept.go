package models

type Dept struct {
	// Model
	ID    uint   `gorm:"primary_key" json:"id" uri:"id"`
	Name  string `json:"name"`
	Users []User `gorm:"foreignkey:DeptID" json:"users"`
}
