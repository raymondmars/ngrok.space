package core

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestClientAuth(t *testing.T) {
	password := "Laojiang$#@897"

	hasher := md5.New()
	hasher.Write([]byte(password))
	encodePwd := hex.EncodeToString(hasher.Sum(nil))

	fmt.Println(encodePwd)
}
