package proxy

import (
	"log"
)

// 代理类实现和被代理对象同样的接口，然后增加一些其他功能，比如增加日志，清理环境等。

type IPlayer interface {
	Start()
	Stop()
}

type player struct {

}

func NewPlayer() *player {
	return &player{}
}

func (p *player) Start() {
	log.Println("Start")
}

func (p *player) Stop() {
	log.Println("Stop")
}

type ProxyPlayer struct {
	player *player
}

func NewProxyPlayer(player *player) *ProxyPlayer {
	return &ProxyPlayer{player}
}

func (p *ProxyPlayer) Start() {
	log.Println("ProxyPlayer banner: open the player")
	p.player.Start()
}

func (p *ProxyPlayer) Stop() {
	
	p.player.Stop()
	log.Println("ProxyPlayer banner: close close the player")
}






