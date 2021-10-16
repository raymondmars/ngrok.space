module raymond.com/ngrok-client

go 1.16

require (
	github.com/gorilla/websocket v1.4.2
	github.com/inconshreveable/mousetrap v1.0.0
	github.com/jteeuwen/go-bindata v3.0.7+incompatible // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/kr/binarydist v0.1.0 // indirect
	github.com/nsf/termbox-go v1.1.1
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/inconshreveable/go-update.v0 v0.0.0-20150814200126-d8b0b1d421aa
	gopkg.in/yaml.v1 v1.0.0-20140924161607-9f9df34309c0
	raymond.com/common v0.0.0
)

replace raymond.com/common => ../common
