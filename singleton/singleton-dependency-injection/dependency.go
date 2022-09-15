package main

import (
	"fmt"
	"sync"
)

/*
Here we are trying to solve the problem with the test and dependency injection in order to use singletons
this is the recommended implementation of singleton
	* sync is important for thread safe
	* lazy is important
	* depend on interfaces not concrete types
*/

func main() {

	cities := []string{"USA", "MX"}

	totalPopulation := GetTotalPopulationEx(GetDatabaseInstance(), cities)
	totalPopulationDummy := GetTotalPopulationEx(&DummyDatabase{}, cities)
	fmt.Println(" population:", totalPopulation)
	fmt.Println("dummy population:", totalPopulationDummy)
}

// Database is an interface that allows us to break the direct dependency on a singleton concrete instance
type Database interface {
	GetPopulation(name string) int
}

type database struct {
	capitals map[string]int
}

func (db *database) GetPopulation(name string) int {
	return db.capitals[name]
}

var once sync.Once
var instance *database

// GetDatabaseInstance still return a concrete real type implementation
func GetDatabaseInstance() *database {
	once.Do(func() {
		instance = populateDatabase()
	})
	return instance
}

// GetTotalPopulationEx is changed in order to depend on a Database which is an interface that represent database
func GetTotalPopulationEx(db Database, cities []string) int {
	total := 0
	for _, city := range cities {
		total += db.GetPopulation(city)
	}
	return total
}

func populateDatabase() *database {
	return &database{capitals: map[string]int{
		"MX":  1000,
		"USA": 2000,
		"CHI": 20000000,
		"JPN": 3333,
	}}
}

// DummyDatabase is only a utility type used for some testing
type DummyDatabase struct {
	dummyData map[string]int
}

func (dd *DummyDatabase) GetPopulation(city string) int {
	return len(city)
}
