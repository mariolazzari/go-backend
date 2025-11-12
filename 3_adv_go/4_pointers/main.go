package main

import (
	"fmt"
	"log"
)

type Truck struct {
	cargo int
}

func main() {
	truckId := 42
	log.Println(truckId)
	log.Println(&truckId)

	anotherTruckId := &truckId
	log.Println(anotherTruckId)
	log.Println(*anotherTruckId)

	truckId = 0
	log.Println(*anotherTruckId)

	t := Truck{}
	fillCargo(&t)
	fmt.Printf("Truck cargo: %d\n", t.cargo)
	fillCargoCopy(t)

	var p *int
	fmt.Println("Pointer default value:", p)
}

func fillCargo(t *Truck) {
	fmt.Printf("t address: %v\n", t)
	t.cargo += 10
}

func fillCargoCopy(t Truck) {
	fmt.Printf("t address: %v\n", t)
	t.cargo += 10
}
