package server

type Options struct {
	HttpAddr   string
	HttpsAddr  string
	TunnelAddr string
	Domain     string
	TlsCrt     string
	TlsKey     string
	Logto      string
	Loglevel   string
}
