// Package system implements methods to manage an elevator control system.
package system

import (
	"fmt"
	"math/rand"
)

// Control is the interface representing the elevator control system.
// It contains calls (i.e. people waiting for an elevator), elevator and the number of floors in the building.
type Control struct {
	calls     []Call
	Elevators []Elevator
	floors    uint8
}

// NewControl returns a new elevator control system.
func NewControl(quantity uint8, floors uint8) (Control, error) {
	if quantity == 0 || quantity > 16 {
		return Control{}, fmt.Errorf("Control only handles up between 1 and 16 elevators.")
	}
	if floors == 0 {
		return Control{}, fmt.Errorf("An elevator needs floors to be interesting.")
	}
	return Control{Elevators: make([]Elevator, quantity), floors: floors}, nil
}

// Pickup creates a call using the given parameter (floor and direction).
// If an elevator in the building is empty, it will come to take the call.
func (c *Control) Pickup(floor uint8, direction Direction) {
	if floor > c.floors {
		return
	}

	call := Call{floor: floor, direction: direction} // A Call object from the parameters given.
	waitingElevatorIndex := -1                       // An index useful if the pickup request cannot be fulfilled directly.

	for i := range c.Elevators {
		if c.Elevators[i].CanTake(call) { // An elevator is available, the client does not wait.
			c.Elevators[i].AddStop(c.randomFloor(c.Elevators[i].Floor())) // Add on which floor the caller wants to go.
			return
		}

		// An elevator is waiting for clients, we save its index in case there is no better elevator.
		if c.Elevators[i].Direction == None {
			waitingElevatorIndex = i
		}
	}

	// The client will have to wait but an elevator is coming if possible.
	if waitingElevatorIndex != -1 {
		c.Elevators[waitingElevatorIndex].AddStop(floor)
		c.Elevators[waitingElevatorIndex].UpdateDirection()
	}

	// We add the call to know that someone is waiting.
	c.calls = append(c.calls, Call{floor: floor, direction: direction})
}

// Step does one step, i.e. moves the elevators once and check the callers and the .
func (c *Control) Step() {
	for i := range c.Elevators {
		c.Elevators[i].Move() // Moves each elevator.

		for j := 0; j < len(c.calls); j++ { // Check if some calls can be taken care of.
			if c.Elevators[i].CanTake(c.calls[j]) {
				c.calls = append(c.calls[:j], c.calls[j+1:]...) // Removing the call from the system.

				c.Elevators[i].AddStop(c.randomFloor(c.Elevators[i].Floor())) // Add a stop to the elevator.
			}
		}

		c.Elevators[i].UpdateDirection()
	}
}

// Steps does a given number of steps.
func (c *Control) Steps(number int) {
	for i := 0; i < number; i++ {
		c.Step()
	}
}

// Print prints the elevator system.
func (c *Control) Print() {
	for i := 0; uint8(i) <= c.floors; i++ {
		floor := c.floors - uint8(i)
		for _, elevator := range c.Elevators {
			fmt.Print("|")
			if elevator.Floor() == floor {
				switch elevator.Direction {
				case Up:
					fmt.Print("▲")
				case None:
					fmt.Print(" ") // Elevator waiting, door "open".
				case Down:
					fmt.Print("▼")
				}
			} else {
				fmt.Print("/") // Elevator on another floor.
			}
			fmt.Print("|")
		}
		floorCalls := 0
		for _, call := range c.calls {
			if call.Floor() == floor {
				floorCalls++
			}
		}
		fmt.Printf(" %d call(s)\n", floorCalls)
	}
}

// randomFloor returns a random uint8 between 0 and max that is not the current floor.
func (c *Control) randomFloor(currentFloor uint8) uint8 {
	randomFloor := currentFloor
	// The client will never ask for the current floor.
	for randomFloor == currentFloor {
		randomFloor = uint8(rand.Intn(int(c.floors + 1))) // + 1 to include the last floor.
	}
	return randomFloor
}
