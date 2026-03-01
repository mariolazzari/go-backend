package main

import (
	"errors"
)

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (*Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks map[string]*Truck
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (t *truckManager) AddTruck(id string, cargo int) error {
	t.trucks["id"] = &Truck{ID: id, Cargo: cargo}
	return nil
}

func (t truckManager) GetTruck(id string) (*Truck, error) {
	truck, ok := t.trucks[id]
	if !ok {
		return nil, ErrTruckNotFound
	}
	return truck, nil
}

func (t *truckManager) RemoveTruck(id string) error {
	delete(t.trucks, id)
	return nil
}

func (t *truckManager) UpdateTruckCargo(id string, cargo int) error {
	t.trucks["id"] = &Truck{ID: id, Cargo: cargo}
	return nil
}
