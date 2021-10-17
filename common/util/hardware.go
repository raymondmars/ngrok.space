package util

import "net"

func GetMacAddress() []string {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	address := []string{}
	for _, item := range interfaces {
		mac := item.HardwareAddr.String()
		if mac != "" {
			address = append(address, mac)
		}
	}
	return address
}
