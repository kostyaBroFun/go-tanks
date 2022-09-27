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
