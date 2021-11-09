package main

import (
	"flag"

	"github.com/ngrok-space/internal/app/server"
)

func ParseArgs() *server.Options {
	httpAddr := flag.String("httpAddr", ":80", "Public address for HTTP connections, empty string to disable")
	httpsAddr := flag.String("httpsAddr", ":443", "Public address listening for HTTPS connections, emptry string to disable")
	tunnelAddr := flag.String("tunnelAddr", ":4443", "Public address listening for ngrok client")
	domain := flag.String("domain", "ngrok.com", "Domain where the tunnels are hosted")
	tlsCrt := flag.String("tlsCrt", "", "Path to a TLS certificate file")
	tlsKey := flag.String("tlsKey", "", "Path to a TLS key file")
	logto := flag.String("log", "stdout", "Write log messages to this file. 'stdout' and 'none' have special meanings")
	loglevel := flag.String("log-level", "DEBUG", "The level of messages to log. One of: DEBUG, INFO, WARNING, ERROR")
	flag.Parse()

	return &server.Options{
		HttpAddr:   *httpAddr,
		HttpsAddr:  *httpsAddr,
		TunnelAddr: *tunnelAddr,
		Domain:     *domain,
		TlsCrt:     *tlsCrt,
		TlsKey:     *tlsKey,
		Logto:      *logto,
		Loglevel:   *loglevel,
	}
}
