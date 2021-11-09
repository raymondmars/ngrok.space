// +build !release,!autoupdate

package client

import "github.com/ngrok-space/internal/app/client/mvc"

// no auto-updating in debug mode
func autoUpdate(state mvc.State, token string) {
}
