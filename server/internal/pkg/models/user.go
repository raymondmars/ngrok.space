package models

import "time"

type User struct {
	ID         int64  `gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	Name       string `gorm:"type:varchar(30)" json:"name"`
	Email      string `gorm:"type:varchar(50)" json:"email" sql:"unique"`
	Password   string `gorm:"type:varchar(50)" json:"password" sql:"index"`
	MacAddress string `gorm:"type:varchar(100)" json:"mac_address"`
	Domains    []Domain
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
