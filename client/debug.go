// +build !release

package main

var (
	rootCrtPaths = []string{"assets/client/tls/ngrokroot.crt", "assets/client/tls/snakeoilca.crt"}
)

func useInsecureSkipVerify() bool {
	return true
}
