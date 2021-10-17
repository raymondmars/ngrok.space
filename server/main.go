package main

import (
	"crypto/tls"
	"math/rand"
	"os"
	"runtime/debug"
	"time"

	"raymond.com/common/msg"
	"raymond.com/ngrok-server/internal/app/core"
	"raymond.com/ngrok-server/internal/pkg/database"

	"raymond.com/common/conn"

	"raymond.com/common/util"

	log "raymond.com/common/log"
)

func NewProxy(pxyConn conn.Conn, regPxy *msg.RegProxy) {
	// fail gracefully if the proxy connection fails to register
	defer func() {
		if r := recover(); r != nil {
			pxyConn.Warn("Failed with error: %v", r)
			pxyConn.Close()
		}
	}()

	// set logging prefix
	pxyConn.SetType("pxy")

	// look up the control connection for this proxy
	pxyConn.Info("Registering new proxy for %s", regPxy.ClientId)
	ctl := core.CommonControlRegistry.Get(regPxy.ClientId)

	if ctl == nil {
		panic("No client found for identifier: " + regPxy.ClientId)
	}

	ctl.RegisterProxy(pxyConn)
}

// Listen for incoming control and proxy connections
// We listen for incoming control and proxy connections on the same port
// for ease of deployment. The hope is that by running on port 443, using
// TLS and running all connections over the same port, we can bust through
// restrictive firewalls.
func tunnelListener(addr string, tlsConfig *tls.Config) {
	// listen for incoming connections
	listener, err := conn.Listen(addr, "tun", tlsConfig)
	if err != nil {
		panic(err)
	}

	log.Info("Listening for control and proxy connections on %s", listener.Addr.String())
	for c := range listener.Conns {
		go func(tunnelConn conn.Conn) {
			// don't crash on panics
			defer func() {
				if r := recover(); r != nil {
					tunnelConn.Info("tunnelListener failed with error %v: %s", r, debug.Stack())
				}
			}()

			tunnelConn.SetReadDeadline(time.Now().Add(core.ConnReadTimeout))
			var rawMsg msg.Message
			if rawMsg, err = msg.ReadMsg(tunnelConn); err != nil {
				tunnelConn.Warn("Failed to read message: %v", err)
				tunnelConn.Close()
				return
			}

			// don't timeout after the initial read, tunnel heartbeating will kill
			// dead connections
			tunnelConn.SetReadDeadline(time.Time{})

			switch m := rawMsg.(type) {
			case *msg.Auth:
				core.NewControl(tunnelConn, m)

			case *msg.RegProxy:
				NewProxy(tunnelConn, m)

			default:
				tunnelConn.Close()
			}
		}(c)
	}
}

func main() {

	database.InstallDb()
	// parse options
	core.OptionParam = core.ParseArgs()

	// init logging
	log.LogTo(core.OptionParam.Logto, core.OptionParam.Loglevel)

	// seed random number generator
	seed, err := util.RandomSeed()
	if err != nil {
		panic(err)
	}
	rand.Seed(seed)

	// init tunnel/control registry
	registryCacheFile := os.Getenv("REGISTRY_CACHE_FILE")
	core.CommonTunnelRegistry = core.NewTunnelRegistry(core.RegistryCacheSize, registryCacheFile)
	core.CommonControlRegistry = core.NewControlRegistry()

	// start listeners
	core.ConnListeners = make(map[string]*conn.Listener)

	// load tls configuration
	tlsConfig, err := core.LoadTLSConfig(core.OptionParam.TlsCrt, core.OptionParam.TlsKey)
	if err != nil {
		panic(err)
	}

	// listen for http
	if core.OptionParam.HttpAddr != "" {
		core.ConnListeners["http"] = core.StartHttpListener(core.OptionParam.HttpAddr, nil)
	}

	// listen for https
	if core.OptionParam.HttpsAddr != "" {
		core.ConnListeners["https"] = core.StartHttpListener(core.OptionParam.HttpsAddr, tlsConfig)
	}

	// ngrok clients
	tunnelListener(core.OptionParam.TunnelAddr, tlsConfig)
}
