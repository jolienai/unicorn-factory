package unicorn

// using in-memory storage for simplicity as in the original

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/exp/slices"
)

type UniCorn struct {
	Name         string
	Capabilities []string
	ProducedAt   time.Time
}

type request struct {
	Amount   int
	Id       int
	Unicorns []UniCorn
}

type UnicornRepository struct {
	mu       sync.RWMutex
	requests []request
	nextID   int
}

func NewUnicornRepository() *UnicornRepository {
	return &UnicornRepository{
		requests: make([]request, 0),
		nextID:   1,
	}
}

func (r *UnicornRepository) Get(requestId int) ([]UniCorn, error) {
	idx := slices.IndexFunc(r.requests, func(c request) bool { return c.Id == requestId })
	if idx < 0 {
		return nil, nil
	}
	return r.requests[idx].Unicorns, nil
}

func (r *UnicornRepository) Create(amount int) (int, error) {
	// returning id as int for simplicity but can be uuid
	r.mu.Lock()
	defer r.mu.Unlock()

	new := request{Amount: amount, Id: r.nextID, Unicorns: make([]UniCorn, 0)}
	r.requests = append(r.requests, new)
	r.nextID++

	fmt.Println("returning id:", new.Id)
	return new.Id, nil
}

func (r *UnicornRepository) Update(amount int, requestId int, ch chan UniCorn) {
	for unicorn := range ch {
		fmt.Printf("Received unicorn: %s for request id: %d \n", unicorn.Name, requestId)
		idx := slices.IndexFunc(r.requests, func(c request) bool { return c.Id == requestId })
		r.requests[idx].Unicorns = append(r.requests[idx].Unicorns, unicorn)
	}
}
