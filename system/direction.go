package system

// Direction is where is going the elevator (up or down)
type Direction int

// Direction is where is going the elevator (up or down)
const (
	Down Direction = -1 + iota
	None
	Up
)
