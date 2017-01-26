package transmission

import (
	"fmt"
	"net/http"
)

const (
	SESSION_HEADER = "X-Transmission-Session-Id"
)

type Client struct {
	Address  string
	Username string
	Password string

	token  string
	client *http.Client
}

func AddressFromHost(host string) string {
	return fmt.Sprintf("http://%s:9091/transmission/rpc", host)
}

func AddressFromHostAndPort(host, port string) string {
	return fmt.Sprintf("http://%s:%s/transmission/rpc", host, port)
}

func AddressFromHostPortAndPath(host, port, path string) string {
	return fmt.Sprintf("http://%s:%s/%s", host, port, path)
}
