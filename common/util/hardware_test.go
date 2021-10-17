package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
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

func TestRandAddress(t *testing.T) {
	// rand.Seed(time.Now().UnixNano())
	seed, err := RandomSeed()
	if err != nil {
		panic(err)
	}
	rand.Seed(seed)

	fmt.Printf("%x\n", rand.Int63())
}
