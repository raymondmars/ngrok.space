package core

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"raymond.com/common/msg"
	"raymond.com/ngrok-server/internal/pkg/database"
	"raymond.com/ngrok-server/internal/pkg/models"
)

type mysqlAuthDb struct {
	Auth *msg.Auth
}

func (ma *mysqlAuthDb) IsValid() bool {
	if ma.Auth.User == "" || ma.Auth.MacAddress == "" {
		return false
	}
	secrets := strings.Split(ma.Auth.User, ":")
	email, password := secrets[0], secrets[1]

	hasher := md5.New()
	hasher.Write([]byte(password))
	encodePwd := hex.EncodeToString(hasher.Sum(nil))
	user := models.User{}
	fmt.Println(email, password, encodePwd)
	database.Db.Where("email = ? and password = ?", email, encodePwd).Preload("Domains").First(&user)
	if user.ID == 0 {
		return false
	}
	if user.MacAddress != "" && user.MacAddress != ma.Auth.MacAddress {
		return false
	}
	if ma.Auth.SubDomains != "" {
		domains := strings.Split(ma.Auth.SubDomains, ",")
		for _, d := range domains {
			exists := false
			for _, e := range user.Domains {
				if d == e.Name {
					exists = true
					break
				}
			}
			if !exists {
				return false
			}
		}
	}
	if user.MacAddress == "" {
		database.Db.Model(&user).Update("mac_address", ma.Auth.MacAddress)
	}

	return true
}
