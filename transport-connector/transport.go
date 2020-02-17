package transport_connector

type transport interface {
	Open(url string) connector
	Close(conn connector) bool

	Send(conn connector, qn string, msg []byte)
	Receive(conn connector, qn string, callback func())
}
