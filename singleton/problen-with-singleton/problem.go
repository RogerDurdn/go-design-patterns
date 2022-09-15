package main

import (
	"fmt"
	"sync"
)

/*
The problem with the singleton is that in some cases can break the dependency inversion
for example in a test case we are using the real database object from the database
if we want to use another one we should add more logic to the singleton
*/

func main() {

	cities := []string{"USA", "MX"}

	totalPopulation := GetToatlPopulation(cities)

	fmt.Println(" population:", totalPopulation)
}

type database struct {
	capitals map[string]int
}

func (db *database) GetPopulation(name string) int {
	return db.capitals[name]
}

var once sync.Once
var instance *database

func GetDatabaseInstance() *database {
	once.Do(func() {
		instance = poblateDatabase()
	})
	return instance
}

func GetToatlPopulation(cities []string) int {
	total := 0
	for _, city := range cities {
		total += GetDatabaseInstance().GetPopulation(city)
	}
	return total
}

func poblateDatabase() *database {
	return &database{capitals: map[string]int{
		"MX":  1000,
		"USA": 2000,
		"CHI": 20000000,
		"JPN": 3333,
	}}
}
