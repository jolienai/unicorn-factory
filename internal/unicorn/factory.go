package unicorn

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type UnicornFactory struct {
	Capabilities []string
	Names        []string
	Adjectives   []string
}

func NewFactory() *UnicornFactory {
	factory := &UnicornFactory{
		Capabilities: make([]string, 0),
		Names:        make([]string, 0),
		Adjectives:   make([]string, 0),
	}
	capabilities, err := GetCapabilities()
	if err != nil {
		panic("please try later, unicorn factory unavailable (capabilities)")
	}
	factory.Capabilities = append(factory.Capabilities, capabilities...)

	names, err := GetNames()
	if err != nil {
		panic("please try later, unicorn factory unavailable (names)")
	}
	factory.Names = append(factory.Names, names...)

	adjectives, err := GetAdjectives()
	if err != nil {
		panic("please try later, unicorn factory unavailable (adjectives)")
	}
	factory.Adjectives = append(factory.Adjectives, adjectives...)

	return factory
}

func (f *UnicornFactory) Produce(amount int, requestId int, ch chan UniCorn) {
	capabilities := f.Capabilities
	names := f.Names
	adj := f.Adjectives

	sleep_time := time.Duration(rand.Intn(1000)) * time.Millisecond

	defer close(ch)

	for j := 1; j < amount+1; j++ {

		fmt.Printf("producing %d items for request id %d \n", j, requestId)
		name := adj[rand.Intn(1345)] + "-" + names[rand.Intn(5800)]
		item := UniCorn{
			Name: name,
		}
		time.Sleep(sleep_time)

		for i := 0; i < 3; i++ {
			cap := capabilities[rand.Intn(17)]
			if !Exists(item.Capabilities, cap) {
				item.Capabilities = append(item.Capabilities, cap)
			}
		}

		item.ProducedAt = time.Now().UTC()

		ch <- item
	}
}

func GetNames() ([]string, error) {
	fn, err := os.Open("petnames.txt")
	if err != nil {
		fmt.Println("please try later, unicorn factory unavailable")
		return []string{}, err
	}
	var names []string
	var scanner = bufio.NewScanner(fn)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}
	return names, nil
}

func GetCapabilities() ([]string, error) {
	return []string{"super strong", "fullfill wishes", "fighting capabilities", "fly", "swim", "sing", "run", "cry", "change color", "talk", "dance", "code", "design", "drive", "walk", "talk chinese", "lazy"}, nil
}

func GetAdjectives() ([]string, error) {
	fa, err := os.Open("adj.txt")
	if err != nil {
		fmt.Println("please try later, unicorn factory unavailable")
		return []string{}, err
	}
	var adj []string
	scanner := bufio.NewScanner(fa)
	for scanner.Scan() {
		adj = append(adj, scanner.Text())
	}
	return adj, nil
}

func Exists(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
