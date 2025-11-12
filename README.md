# Backend Engineering with Go

## Introduction

### Resources

- [Github](https://github.com/sikozonpc/GopherSocial)
- [Context](https://youtu.be/Q0BdETrs1Ok)
- [Error Handling](https://youtu.be/dKUiCF3abHc)
- [Interfaces](https://youtu.be/4OVJ-ir9hL8?si=nZcSoQrTXrYh69y4)
- [Maps](https://youtu.be/999h-iyp4Hw?si=fPLtWRs7DWIVBIk-)
- [Pointers](https://youtu.be/DVNOP1LE3Mg?si=KXaKeHeIipjLg1HZ)
- [Goroutines & Channels](https://youtu.be/3QESpVGiiB8?si=kqpETtKp73Abyiyw)

## Project architecture

### Design principles for REST API

- [12 factors](https://12factor.net/)
- [Roy Fielding](https://ics.uci.edu/~fielding/pubs/dissertation/fielding_dissertation.pdf)
- [Richardson Maturity Model](https://martinfowler.com/articles/richardsonMaturityModel.html)

## Advanced Go

### Effective error handling

```go
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
```

### Interfaces

```go
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
```

### Tests

```sh
go mod init github.com/mariolazzari/go-backend
go test -v *.go
```

```go
package main

import (
	"testing"
)

func MainTest(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {
		t.Run("should load and unload a truck", func(t *testing.T) {
			nt := &NormalTruck{id: "T1", cargo: 10}
			et := &ElectrictTruck{id: "eT2", cargo: 20}

			err := processTruck(nt)
			if err != nil {
				t.Fatalf("Error processing truck: %s", err)
			}

			err = processTruck(et)
			if err != nil {
				t.Fatalf("Error processing truck: %s", err)
			}

			// asserting
			if nt.cargo != 0 {
				t.Fatal("Normal cargo should be 0")
			}
		})
	})
}
```

### Pointers

```go
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
```

### Goroutines

```go

```
