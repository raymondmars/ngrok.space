package models

import "time"

type User struct {
	ID         int64  `gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	Name       string `gorm:"type:varchar(30)" json:"name"`
	Email      string `gorm:"type:varchar(50);uniqueIndex" json:"email"`
	Password   string `gorm:"type:varchar(50);index" json:"password"`
	AuthToken  string `gorm:"type:varchar(60);index" json:"auth_token"`
	MacAddress string `gorm:"type:varchar(100)" json:"mac_address"`
	Domains    []Domain
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
