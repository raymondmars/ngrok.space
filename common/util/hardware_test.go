package util

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"math/rand"
	"testing"

	"github.com/rs/xid"
	"github.com/segmentio/ksuid"
)

func TestGetMacAddress(t *testing.T) {
	mac := GetMacAddress()
	fmt.Println(mac)
}

func TestClientAuth(t *testing.T) {
	// email := "bydanta@163.com"

	// hasher := md5.New()
	// hasher.Write([]byte(email))
	// encodEmail := hex.EncodeToString(hasher.Sum(nil))

	a := xid.New()
	b := ksuid.New()
	token := fmt.Sprintf("%s_%s", a, b)
	fmt.Println(token)

	// fmt.Println("abC" == "abc")
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

func Sha1(s string) string {
	sum := sha1.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}
