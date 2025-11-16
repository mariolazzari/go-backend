# Backend Engineering with Go

## Introduction

### Resources

- [Project](https://github.com/sikozonpc/GopherSocial/tree/main)
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
package main

import (
	"fmt"
	"log"
	"os"
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
	// Wait group
	var wg sync.WaitGroup
	// wg.Add(len(trucks))

	for _, truck := range trucks {
		wg.Add(1)
		go func(t Truck) {
			if err := processTruck(truck); err != nil {
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
```

### Context and timeouts

```go
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
```

### Concurrency with channels

```go
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
```

### Maps

```go
package main

func main() {
	m := make(map[string]int)
	// m["a"] = 1
	// a, exists := m["a"]

	if _, ok := m["a"]; ok {
		// ...
	}

	delete(m, "a")
	clear(m)
}
```

### Capstone project

```go
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
```

## From TCP to HTTP

### TCP server

[net package](https://pkg.go.dev/net)

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {

			log.Fatal(err)
			return
		}

		// handle multiple connection
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// create new reader from connection
	reader := bufio.NewReader(conn)
	// read from clinet command line
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(conn, "Error reading command: %v\n", err)
		return
	}

	parts := strings.SplitN(strings.TrimSpace(line), " ", 2)
	if len(parts) != 2 {
		fmt.Fprintf(conn, "Invalid command\n")
		return
	}
}
```

### Understanding routing

```go
package route

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// connecto to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// write to server
	fmt.Fprintf(conn, "GET /index.html\n")

	// read server response
	bs := make([]byte, 1024)
	n, err := conn.Read(bs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs[:n]))
}
```

### HTTP server

[http](https://pkg.go.dev/net/http)

```go
package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ciao Mario"))
}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	api := &api{addr: ":8080"}

	// init mux
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	server.ListenAndServe()
}
```

### JSON requests

```go
package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type api struct {
	addr string
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var users = []User{}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// encode users slice to json
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}
	err = insertUser(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func insertUser(u User) error {
	if u.FirstName == "" {
		return errors.New("First name is mandatory")
	}
	if u.LastName == "" {
		return errors.New("Last name is mandatory")
	}

	// storage validation
	for _, user := range users {
		if user.FirstName == u.FirstName && user.LastName == u.LastName {
			return errors.New("User already saved")
		}
	}
	users = append(users, u)

	return nil
}

func main() {
	api := &api{addr: ":8080"}

	// init mux
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```

## Scaffolding API server

### Development enviroment

```sh
mkdir social
cd social
go mod init github.com/mariolazzari/go-backend/social
mkdir bin
mkdir cmd
cd cmd
mkdir api
mkdir cmd
cd migrate
mkdir migrations
cd ../../social
mkdir iternal
mkdir docs
mkdir scripts
mkdir web
```

#### Links

[Docker](https://www.docker.com/products/docker-desktop/)
[Course](https://www.youtube.com/watch?v=7VLmLOiQ3ck&t=3647s)
[Course](https://www.youtube.com/watch?v=h3fqD6IprIA)
[Course](https://www.youtube.com/watch?v=s3XItrqfccw)

### Clean layered architecture

- Separation of concerns: each layer must be separated by a clear barrier
  - transport
  - service
  - storage
- Dependency inversion principle: inject dependency in each layer, do not call them directly, promoting loose coupling and testing
- Adaptibilty to change: modularity means flexibility
-

#### Links

[Book](https://www.amazon.com/Clean-Architecture-Craftsmans-Software-Structure/dp/0134494164)

### HTTP server and API

```sh
go get -u github.com/go-chi/chi/v5
```

### Hor reload

[Air](https://github.com/air-verse/air)

```sh
go install github.com/air-verse/air@latest
air init
air
```

### Enviroment varaibles

[direnv](https://direnv.net/)
[godotenv](https://github.com/joho/godotenv)
[Factor config](https://12factor.net/config)

## Databases

### Repository pattern

[Repository pattern in Go](https://threedots.tech/post/repository-pattern-in-go/)

```sh
go install github.com/lib/pq@latest
```

```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store interface {
	GetByID(id int) (*User, error)
}

type application struct {
	store Store
}

type User struct {
	ID   int
	Name string
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetByID(id int) (*User, error) {
	row := r.db.QueryRow("select id, name from users where id = $1", id)

	var user User
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func main() {
	connStr := "user=username dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepository := NewPostgresUserRepository(db)

	user, err := userRepository.GetByID(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User: %+v", user)
	}

}

type InMemoryRepository struct {
	users []User
}

func (r *InMemoryRepository) GetByID(id int) (*User, error) {
	return nil, nil
}
```

### Implementing repository pattern

```go

```
