// +build release

package main

var (
	rootCrtPaths = []string{"assets/client/tls/ngrokroot.crt"}
)

func useInsecureSkipVerify() bool {
	return false
}
