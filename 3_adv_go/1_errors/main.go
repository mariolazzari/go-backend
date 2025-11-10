package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("Truck not implemented")
	ErrNotFound       = errors.New("Truck not found")
)

type Truck struct {
	id string
}

func (t *Truck) LoadCargo() error {
	return ErrNotFound
}

// loading and unloading
func processTruck(t Truck) error {
	fmt.Printf("Processing truck: %s\n", t.id)

	if err := t.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cardo: %w", err)
	}

	return ErrNotImplemented
}

func main() {
	trucks := []Truck{{id: "Truck-1"}, {id: "Truck-2"}, {id: "Truck-3"}}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived\n", truck.id)

		// err := processTruck(truck)
		// if err != nil {
		// 	log.Fatalf("Error processing truck: %s", err)
		// }

		if err := processTruck(truck); err != nil {
			switch err {
			case ErrNotImplemented:
			case ErrNotFound:
			}
			if errors.Is(err, ErrNotImplemented) {

			}

			if errors.Is(err, ErrNotFound) {

			}

			log.Fatalf("Error processing truck: %s", err)
		}
	}
}
