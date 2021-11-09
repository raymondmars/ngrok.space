package core

import (
	"strings"

	"github.com/ngrok-space/internal/app/server/database"
	"github.com/ngrok-space/internal/app/server/models"
	"github.com/ngrok-space/internal/pkg/msg"
)

type mysqlAuthDb struct {
	Auth *msg.Auth
}

func (ma *mysqlAuthDb) IsValid() bool {
	if ma.Auth.User == "" || ma.Auth.MacAddress == "" {
		return false
	}
	// secrets := strings.Split(ma.Auth.User, ":")
	// email, password := secrets[0], secrets[1]

	// hasher := md5.New()
	// hasher.Write([]byte(password))
	// encodePwd := hex.EncodeToString(hasher.Sum(nil))
	user := models.User{}
	database.Db.Where("auth_token = ?", ma.Auth.User).Preload("Domains").First(&user)
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
