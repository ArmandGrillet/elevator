# Elevator

## Features and differences with the given Scala interface

Features of the control system:
- Querying the state of the elevators (what floor are they on and where they are going),
- Receiving a pickup request,
- Time-stepping the simulation.

Differences compared to the given interface:
- Status() has been replaced by Control.Elevators (the only public field of the Control structure).
- Update(int, int, int) does not exist (no time to implement it).
- Print() is an added method to print the control system state, used in `main.go`.

## Structure of the project
There is a `main.go` to show how to use the `system` package containing the interface. The package contains:
- `control.go`, the control system with its calls (people waiting), elevators and a private number of floors.
- `elevator.go` representing an elevator with a public direction, a readable floor and a private list of stops.
- `call.go` and `direction.go`, two basic structs representing calls and directions (up, none or down).
- `control_test.go` and `elevator_test.go`, two files containing unit tests.

## The scheduling
The most important thing for an elevator is its direction. If it goes up, it will fulfill all the stops above its current floor before going down even if someone asked for going down before other stops that are up. This logic can be found in `elevator.UpdateDirection()`, a method used when someone picks up and when a step is done.

## How to run the project
```
$ go get github.com/ArmandGrillet/elevator
$ cd $GOPATH/src/github.com/ArmandGrillet/elevator
$ go run main.go
```

## How to test the project
```
$ go get github.com/ArmandGrillet/elevator
$ cd $GOPATH/src/github.com/ArmandGrillet/elevator/system
$ go test
```

## How to read the documentation
```
$ cd $GOPATH/src/github.com/ArmandGrillet/elevator/system
$ godoc ./
```

## How to use the project
Import `"github.com/ArmandGrillet/elevator/system"` in your Go project.
