package util

import (
	"fmt"
	"testing"
)

func TestGetMacAddress(t *testing.T) {
	mac := GetMacAddress()
	fmt.Println(mac)
}
