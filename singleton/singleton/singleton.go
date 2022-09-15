package main

import (
	"fmt"
	"sync"
)

/*
Provide of a unique instance that has access globally by the system
Used in cost objects
Mostly used with lazy initialization
*/

func main() {

	db := GetDatabaseInstance()

	usaPopulation := db.GetPopulation("USA")
	fmt.Println("usa population:", usaPopulation)
}

type database struct {
	capitals map[string]int
}

func (db *database) GetPopulation(name string) int {
	return db.capitals[name]
}

/*
The correct way to use a singleton on go is make sure that is thread safe
option 1: sync.Once
option 2: init()
another thing in mind is the laziness initialization
to have laziness we are going to use the sync.Once approach
*/

var once sync.Once
var instance *database

func GetDatabaseInstance() *database {
	once.Do(func() {
		instance = poblateDatabase()
	})
	return instance
}

func poblateDatabase() *database {
	return &database{capitals: map[string]int{
		"MX":  1000,
		"USA": 2000,
		"CHI": 20000000,
		"JPN": 3333,
	}}
}
