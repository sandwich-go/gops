package gops

import (
	"net"
	"testing"

	"github.com/sandwich-go/gops/signal"
)

func TestCmd(t *testing.T) {
	bytes, err := cmd(net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8991}, signal.StackTrace)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bytes))
}
