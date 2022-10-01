package game

type Cmd int16

const (
	NoCommand Cmd = iota
	GoForward
	GoBack
	RotateLeft
	RotateRight
	Fire
)

type Player struct {
	id   int64
	cmds <-chan Cmd
}

func (p *Player) SendCommand(cmd Cmd) {

}

func (p *Player) SubscribeOnCommands() <-chan Cmd {
	panic("not implemented")
}

func (p *Player) Unsubscribe() {

}
