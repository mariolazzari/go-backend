package main

import (
	"fmt"
	"log"
	"sync"
	"time"
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
	t.cargo += 10
	t.battery += 10
	return nil
}

func (t *ElectrictTruck) UnloadCargo() error {
	t.cargo = 0
	t.battery -= 5
	return nil
}

func processTruck(truck Truck) error {
	fmt.Printf("Start processing truck: %+v\n", truck)

	// simulate delay
	time.Sleep(time.Second)

	err := truck.LoadCargo()
	if err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	err = truck.UnloadCargo()
	if err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	fmt.Printf("Finish processing truck: %+v\n", truck)
	return nil
}

func processFleet(trucks []Truck) error {
	var wg sync.WaitGroup
	var errorsChan = make(chan error, len(trucks))

	for _, truck := range trucks {
		wg.Add(1)
		go func(t Truck) {
			if err := processTruck(truck); err != nil {
				log.Panicln(err)
				errorsChan <- err
			}
			wg.Done()
		}(truck)
	}

	wg.Wait()
	defer close(errorsChan)

	select {
	case err := <-errorsChan:
		return err
	default:
		return nil
	}
}

func main() {
	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectrictTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectrictTruck{id: "ET2", cargo: 0, battery: 50},
	}

	if err := processFleet(fleet); err != nil {
		fmt.Printf("Error processing fleet: %v\n", err)
		return
	}

	fmt.Println("All trucks processed")
}
