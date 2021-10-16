module raymond.com/ngrok-server

go 1.16

require (
	github.com/inconshreveable/go-vhost v0.0.0-20160627193104-06d84117953b
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475
	gorm.io/gorm v1.21.16
	raymond.com/common v0.0.0
)

replace raymond.com/common => ../common
