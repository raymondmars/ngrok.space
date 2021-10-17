package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGetMacAddress(t *testing.T) {
	mac := GetMacAddress()
	fmt.Println(mac)
}

func TestClientAuth(t *testing.T) {
	password := ""

	hasher := md5.New()
	hasher.Write([]byte(password))
	encodePwd := hex.EncodeToString(hasher.Sum(nil))

	fmt.Println(encodePwd)
}
