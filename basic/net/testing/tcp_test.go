package testing

import (
	"basic/net/tcp"
	"testing"
)

func TestTcpServer(t *testing.T) {
	tcp.ServerStart()
}

func TestTcpClient(t *testing.T) {
	tcp.ClientStart()
}
