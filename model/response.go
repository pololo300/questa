package Models

import "gorm.io/gorm"

type Response struct {
	gorm.Model

	UserID uint `json:"user_id" gorm:"index;not null;uniqueIndex:uniq_user_poll"`
	User   User `json:"user" `

	PollID uint `json:"poll_id" gorm:"index;not null;uniqueIndex:uniq_user_poll"`
	Poll   Poll `json:"poll"`

	OptionID uint   `json:"option_id" gorm:"index;not null"`
	Option   Option `json:"option"`
}
