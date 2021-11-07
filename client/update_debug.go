// +build !release,!autoupdate

package main

import (
	"ngrok.space/ngrok-client/mvc"
)

// no auto-updating in debug mode
func autoUpdate(state mvc.State, token string) {
}
