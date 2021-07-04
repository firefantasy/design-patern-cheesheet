package command

import (
	"log"
)

// 命令模式将某个对象的方法单独封装成一个个命令类，这样可以灵活组合这些命令类。
// 甚至可以解决某些语言不能传递函数的问题。
// 命令模式可以用于批处理，任务队列，undo，redo，排队，异步，给命令执行记录日志

type Command interface {
	Execute()
}

type PlayCommand struct {
	*player
}

func NewPlayCommnd(player *player) *PlayCommand {
	return &PlayCommand{ player }
}

func (p *PlayCommand) Execute() {
	p.play()
}

type StopCommand struct {
	*player
}

func NewStopCommnd(player *player) *StopCommand {
	return &StopCommand{ player }
}

func (p *StopCommand) Execute() {
	p.stop()
}

type NextCommand struct {
	*player
}

func NewNextCommnd(player *player) *NextCommand {
	return &NextCommand{ player }
}

func (p *NextCommand) Execute() {
	p.next()
}

type PreviousCommand struct {
	*player
}

func NewPreviousCommnd(player *player) *PreviousCommand {
	return &PreviousCommand{ player } 
}

func (p *PreviousCommand) Execute() {
	p.previous()
}


type player struct {

}

func NewPlayer() *player {
	return &player{}
}

func (player) play() {
	log.Println("start play")
}

func (player) stop() {
	log.Println("stop play")
}

func (player) next() {
	log.Println("switch next song")
}

func (player) previous() {
	log.Println("switch previous song")
}
