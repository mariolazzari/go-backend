package main

import (
	"errors"
	"fmt"
)

var (
	ErrNotImplemented = errors.New("Truck not implemented")
	ErrNotFound       = errors.New("Truck not found")
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 10
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

type ElectrictTruck struct {
	id      string
	cargo   int
	battery float64
}

func (t *ElectrictTruck) LoadCargo() error {
	return nil
}

func (t *ElectrictTruck) UnloadCargo() error {
	return nil
}

// loading and unloading
func processTruck(t Truck) error {
	if err := t.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}
	fmt.Printf("Processing truck %+v\n", t)

	return nil
}

func main() {
	trucks := []NormalTruck{{id: "Truck-1", cargo: 10}, {id: "Truck-2", cargo: 20}, {id: "Truck-3"}}
	for _, truck := range trucks {
		processTruck(&truck)
	}

	eTrucks := []ElectrictTruck{{id: "Truck-1", cargo: 10, battery: 100}, {id: "Truck-2", cargo: 20, battery: 200}, {id: "Truck-3", cargo: 30}}
	for _, truck := range eTrucks {
		processTruck(&truck)
	}

}
