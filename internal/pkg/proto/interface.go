package proto

import "github.com/ngrok-space/internal/pkg/conn"

type Protocol interface {
	GetName() string
	WrapConn(conn.Conn, interface{}) conn.Conn
}
