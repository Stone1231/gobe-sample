package models

import "time"

type Model struct {
	ID        uint      `gorm:"primary_key" json:"id" uri:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
