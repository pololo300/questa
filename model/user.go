package Models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string     `json:"name" gorm:"not null;index"`
	CreatedPolls []Poll     `json:"created_polls" gorm:"foreignKey:CreatedByID;"`
	Responses    []Response `json:"responses" gorm:"foreignKey:UserID"`
}
