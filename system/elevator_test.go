package system

import (
	"reflect"
	"testing"
)

func TestAddStop(t *testing.T) {
	e := Elevator{}
	e.AddStop(4)
	if !reflect.DeepEqual(e.stops, []uint8{4}) {
		t.Error("Elevator{}.AddStop(4) != Elevator{stops: []uint8{4}}")
	}
	e.AddStop(4)
	if !reflect.DeepEqual(e.stops, []uint8{4}) {
		t.Error("Elevator{stops: []uint8{4}}.AddStop(4) != Elevator{stops: []uint8{4}}")
	}
	e.AddStop(2)
	if !reflect.DeepEqual(e.stops, []uint8{4, 2}) {
		t.Error("Elevator{stops: []uint8{4}}.AddStop(4, 2) != Elevator{stops: []uint8{4, 2}}")
	}
}

func TestCanTake(t *testing.T) {
	e := Elevator{floor: 2, Direction: Up, stops: []uint8{3}}
	upCall := Call{floor: 2, direction: Up}
	downCall := Call{floor: 2, direction: Down}
	wrongFloorCall := Call{floor: 4, direction: Down}

	if !e.CanTake(upCall) {
		t.Error("Elevator{Floor: 2, Direction: Up}.CanTake(Call{floor: 2, direction: Up}) = false")
	}
	if e.CanTake(wrongFloorCall) {
		t.Error("Elevator{Floor: 2, Direction: Up}.CanTake(Call{floor: 2, direction: Down}) = true")
	}
	if e.CanTake(downCall) {
		t.Error("Elevator{Floor: 2, Direction: Up}.CanTake(Call{floor: 4, direction: Down}) = true")
	}
}

func TestMove(t *testing.T) {
	e := Elevator{floor: 3, Direction: Up, stops: []uint8{4, 2}}
	e.Move()
	if !reflect.DeepEqual(e, Elevator{floor: 4, Direction: Up, stops: []uint8{2}}) {
		t.Error("Elevator{floor: 3, Direction: Up, stops: []uint8{4, 2}}.Move() != Elevator{floor: 4, Direction: Up, stops: []uint8{2}}")
	}

	e = Elevator{floor: 4, Direction: Down, stops: []uint8{2}}
	e.Move()
	if !reflect.DeepEqual(e, Elevator{floor: 3, Direction: Down, stops: []uint8{2}}) {
		t.Error("Elevator{floor: 4, Direction: Up, stops: []uint8{2}}.Move() != Elevator{floor: 3, Direction: Down, stops: []uint8{2}}")
	}

	e.Move()
	if !reflect.DeepEqual(e, Elevator{floor: 2, Direction: Down, stops: []uint8{}}) {
		t.Error("Elevator{floor: 2, Direction: Up, stops: []uint8{2}}.Step() != Elevator{floor: 2, Direction: none, stops: []uint8{}}")
	}
}

func TestUpdateDirection(t *testing.T) {
	e := Elevator{floor: 4, Direction: Up, stops: []uint8{2}}
	e.UpdateDirection()
	if !reflect.DeepEqual(e, Elevator{floor: 4, Direction: Down, stops: []uint8{2}}) {
		t.Error("Elevator{floor: 4, Direction: Up, stops: []uint8{2}}.UpdateDirection() != Elevator{floor: 4, Direction: Down, stops: []uint8{2}}")
	}
	e.Move()
	e.Move()
	e.UpdateDirection()
	if !reflect.DeepEqual(e, Elevator{floor: 2, Direction: None, stops: []uint8{}}) {
		t.Error("UpdateDirection does not set the elevator direction to none when it has nothing to do")
	}
}
