package core

import (
	"time"

	"ngrok.space/common/conn"
)

const (
	RegistryCacheSize uint64        = 1024 * 1024 // 1 MB
	ConnReadTimeout   time.Duration = 10 * time.Second
)

// GLOBALS
var (
	CommonTunnelRegistry  *TunnelRegistry
	CommonControlRegistry *ControlRegistry

	// XXX: kill these global variables - they're only used in tunnel.go for constructing forwarding URLs
	OptionParam   *Options
	ConnListeners map[string]*conn.Listener
)
