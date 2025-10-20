package Models

import "gorm.io/gorm"

type PollType string

const (
	PollSingle   PollType = "single"
	PollMultiple PollType = "multiple"
	PollText     PollType = "text"
	PollRange    PollType = "range"
)

type Poll struct {
	gorm.Model
	CreatedByID uint     `json:"created_by_id" gorm:"index"`
	CreatedBy   User     `json:"created_by"    gorm:"foreignKey:CreatedByID;"`
	Type        PollType `json:"type" gorm:"type:varchar(16);index"`

	Title string `json:"title" gorm:"not null"`
	Body  string `json:"body" gorm:"type:text"`

	Options []Option `json:"options"   gorm:"foreignKey:PollID;"`

	Responses []Response `json:"responses" gorm:"foreignKey:PollID;"`
}
