package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type contextKey string

var UserIDKey contextKey = "userID"

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

func processTruck(ctx context.Context, truck Truck) error {
	fmt.Printf("Start processing truck: %+v\n", truck)

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	// simulate delay
	delay := time.Second * 3
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(delay):
		break
	}

	// access user id
	userId := ctx.Value(UserIDKey)
	log.Println(userId)

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

func processFleet(ctx context.Context, trucks []Truck) error {
	// Wait group
	var wg sync.WaitGroup
	// wg.Add(len(trucks))

	for _, truck := range trucks {
		wg.Add(1)
		go func(t Truck) {
			if err := processTruck(ctx, truck); err != nil {
				log.Panicln(err)
				os.Exit(1)
			}
			wg.Done()
		}(truck)
	}

	wg.Wait()

	return nil
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, UserIDKey, 42)

	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectrictTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectrictTruck{id: "ET2", cargo: 0, battery: 50},
	}

	if err := processFleet(ctx, fleet); err != nil {
		fmt.Printf("Error processing fleet: %v\n", err)
		return
	}

	fmt.Println("All trucks processed")
}
