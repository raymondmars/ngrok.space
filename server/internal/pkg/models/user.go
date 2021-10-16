package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	MacAddress string `json:"mac_address"`
	Domain     string `json:"domain"`
}
