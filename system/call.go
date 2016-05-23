package system

// Call is the structure representing a pick-up request for the control system.
type Call struct {
	floor     uint8
	direction Direction
}

// Floor returns on which floor the call has been made.
func (c *Call) Floor() uint8 {
	return c.floor
}

// Direction returns on which direction wants to go the caller.
func (c *Call) Direction() Direction {
	return c.direction
}
