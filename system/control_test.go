package system

import (
	"reflect"
	"testing"
)

func TestNewControl(t *testing.T) {
	_, err := NewControl(0, 0)
	if err == nil {
		t.Error("NewControl(0, 0) does not return an error")
	}

	_, err = NewControl(42, 42)
	if err == nil {
		t.Error("NewControl(42, 42) does not return an error")
	}

	_, err = NewControl(42, 7)
	if err == nil {
		t.Error("NewControl(42, 7) does not return an error")
	}

	c, err := NewControl(16, 7)
	if err != nil {
		t.Error("NewControl(16, 7) returns an error")
	}

	if len(c.Elevators) != 16 {
		t.Error("NewControl(16, 7) does not return a system with 16 elevators")
	}

	if c.floors != 7 {
		t.Error("NewControl(16, 7) does not return a system with 7 floors")
	}
}

func TestPickup(t *testing.T) {
	c, _ := NewControl(16, 7)

	c.Pickup(0, Up)
	if len(c.calls) != 0 {
		t.Error("NewControl(16, 7).Pickup(0, Up) creates a call but the first elevator should take the call directly")
	}
	if len(c.Elevators[0].stops) != 1 {
		t.Error("NewControl(16, 7).Pickup(0, Up) does not give the call to the first available elevator")
	}

	c.Pickup(8, Up)
	if len(c.calls) != 0 {
		t.Error("NewControl(16, 7).Pickup(8, Up) worked but there is no 8th floor")
	}

	c.Pickup(7, Down)
	if !reflect.DeepEqual(c.calls, []Call{Call{7, Down}}) {
		t.Error("NewControl(16, 7).Pickup(7, Up) does not add a call")
	}

	c.Elevators = []Elevator{Elevator{floor: 4, Direction: Up}, Elevator{floor: 2, Direction: None}}
	c.Pickup(3, Up)
	if len(c.Elevators[1].stops) != 1 {
		t.Error("Pickup request but no elevator is coming")
	}
	if len(c.calls) != 2 {
		t.Error("Pickup request waiting but not added to c.calls")
	}

	c.Pickup(1, Up)
	if len(c.calls) != 3 {
		t.Error("Pickup request waiting but not added to c.calls")
	}
}

func TestStep(t *testing.T) {
	c, _ := NewControl(2, 7)
	c.Pickup(0, Up)
	c.Step()
	if c.Elevators[0].Direction != Up && c.Elevators[0].Floor() != 1 {
		t.Error("Step with an elevator that needs to set but the elevator did not move")
	}
}

func TestRandomFloor(t *testing.T) {
	c, _ := NewControl(2, 7)
	for i := 0; i < 1337; i++ {
		if c.randomFloor(4) == 4 {
			t.Error("randomFloor returns the currentFloor")
		}
	}
}
