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

```
