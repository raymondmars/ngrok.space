package core

import (
	"time"

	"github.com/ngrok-space/internal/app/server"
	"github.com/ngrok-space/internal/pkg/conn"
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
	OptionParam   *server.Options
	ConnListeners map[string]*conn.Listener
)
