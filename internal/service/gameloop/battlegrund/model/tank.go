package model

type Tank struct {
}

type MoveType int8

const (
	MoveForward MoveType = iota
	MoveBack
)

func (t *Tank) Move(moveType MoveType) {

}

func (t *Tank) Rotate() {

}

func (t *Tank) Fire() {

}
