package biz

type Direction int

const (
	Up Direction = iota
	Left
	Down
	Right
)

var (
	DirectionString = []string{
		"Up",
		"Left",
		"Down",
		"Right",
	}
)
