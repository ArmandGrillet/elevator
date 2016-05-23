package system

// Elevator is the basic structure representing an elevator.
type Elevator struct {
	floor     uint8
	Direction Direction
	stops     []uint8
}

// Floor returns the floor of the elevator.
func (e *Elevator) Floor() uint8 {
	return e.floor
}

// AddStop adds a stop to the elevator, it does not affect its direction (see CheckDirection for this).
func (e *Elevator) AddStop(floor uint8) {
	if e.indexOf(floor) == -1 {
		e.stops = append(e.stops, floor)
	}
}

// CanTake returns if an elevator can take a new client.
func (e *Elevator) CanTake(c Call) bool {
	// If the elevator is on the same floor and can go where the caller wants, the caller enters.
	if e.floor == c.floor && (e.Direction == c.direction || len(e.stops) == 0) {
		return true
	}
	return false
}

// Move makes the elevator do one step and removes the new floor from its stops.
func (e *Elevator) Move() {
	if e.Direction == Up {
		e.floor++
	} else if e.Direction == Down {
		e.floor--
	}

	floorIndex := e.indexOf(e.floor)
	if floorIndex != -1 {
		// Removes the current floor if it was one of the stop.
		e.stops = append(e.stops[:floorIndex], e.stops[floorIndex+1:]...)
	}
}

// UpdateDirection updates the direction of the elevator if needed.
func (e *Elevator) UpdateDirection() {
	if len(e.stops) == 0 {
		e.Direction = None
		return
	}

	if e.Direction == None { // The elevator was not moving, it now goes wherever it has to.
		if e.floor < e.stops[0] {
			e.Direction = Up
		} else if e.floor > e.stops[0] {
			e.Direction = Down
		}
	} else {
		// Compute the minimum and maximum floors where needs to go the elevator.
		minFloor, maxFloor := e.floor, e.floor
		for _, stop := range e.stops {
			if stop > maxFloor {
				maxFloor = stop
			}
			if stop < minFloor {
				minFloor = stop
			}
		}

		// If the elevator has only clients going the other way, it changes its direction.
		if e.Direction == Up && maxFloor <= e.Floor() {
			e.Direction = Down
		} else if e.Direction == Down && minFloor >= e.Floor() {
			e.Direction = Up
		}
	}
}

// Return the position of a given floor in the list of stops of the elevator.
func (e *Elevator) indexOf(floor uint8) int {
	for i, stop := range e.stops {
		if stop == floor {
			return i
		}
	}
	return -1
}
