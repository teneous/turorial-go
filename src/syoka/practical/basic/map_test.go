//trade-off
package basic

import (
	"fmt"
	"testing"
)

func TestTradeOffMap(t *testing.T) {
	cities := []City{
		NewCity("Swore", "a", 1000),
		NewCity("Alex", "b", 1000),
		NewCity("Swore", "c", 1000),
		NewCity("Axc", "e", 500),
		NewCity("Axc", "f", 60),
	}

	tradeoff(cities)
}

func tradeoff(cities []City) {
	cityMap := make(map[string][]City)
	for _, entry := range cities {
		if cityMap[entry.province] != nil {
			cities := cityMap[entry.province]
			cityMap[entry.province] = append(cities, entry)
		} else {
			cityMap[entry.province] = []City{entry}
		}
	}
	PRINT(cityMap)
}

type City struct {
	province   string
	name       string
	population int64
}

func NewCity(province string, name string, pop int64) City {
	return City{
		province:   province,
		name:       name,
		population: pop,
	}
}

func PRINT(data map[string][]City) {
	for k, v := range data {
		fmt.Printf("Province: %s , cities: %v \n", k, v)
	}
}
