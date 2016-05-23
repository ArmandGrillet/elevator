package main

import (
	"fmt"
	"log"

	"github.com/ArmandGrillet/elevator/system"
)

func main() {
	controlSystem, err := system.NewControl(2, 5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("This is the elevator when we start. Both elevators are at floor 0, ready to serve:")
	controlSystem.Print()

	fmt.Println("Pickup request floor 4, the last free elevator starts moving:")
	controlSystem.Pickup(4, system.Down)
	controlSystem.Print()

	fmt.Println("3 steps later:")
	controlSystem.Steps(3)
	controlSystem.Print()

	fmt.Println("Pickup request floor 2, the first elevator moves after one step:")
	controlSystem.Pickup(2, system.Up)
	controlSystem.Step()
	controlSystem.Print()
}
