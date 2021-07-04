package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	player := NewProxyPlayer(NewPlayer())
	player.Start()
	player.Stop()
}