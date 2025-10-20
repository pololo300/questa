import "gorm.io/gorm"

type Option struct {
	gorm.Model
	PollID uint   `json:"poll_id" gorm:"index;not null"`
	Poll   Poll   `json:"-"`
	Label  string `json:"label" gorm:"not null"`
	Value  string `json:"value" gorm:"not null;index"`
	Order  int    `json:"order" gorm:"index"`
	Active bool   `json:"active" gorm:"default:true"`
}
