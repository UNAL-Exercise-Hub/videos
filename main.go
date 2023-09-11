package main

import (
	"github.com/UNWorkout/Videos/orm"
	"github.com/UNWorkout/Videos/routes"
)

func main() {
	orm.Connection()
	routes.Start()
}
