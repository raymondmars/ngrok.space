package models

type Domain struct {
	ID     int64  `gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	UserID int64  `json:"user_id"`
	Name   string `gorm:"type:varchar(80)" json:"name"`
}
