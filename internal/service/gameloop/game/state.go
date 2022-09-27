package game

type ObjectType int16

const (
	EmptyObject ObjectType = iota
	Tank
	Projectile
	Wall
)

type Coordinate struct {
	x int
	y int
}

type Object struct {
	objectType ObjectType
	coordinate Coordinate
}

type State struct {
	objects []Object
}
