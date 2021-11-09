package client

type Options struct {
	Config    string
	Logto     string
	Loglevel  string
	Authtoken string
	Httpauth  string
	Hostname  string
	Protocol  string
	Subdomain string
	Command   string
	Args      []string
}
