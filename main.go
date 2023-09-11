package main

import (
	"github.com/jbautistas/videos_ms/orm"
	"github.com/jbautistas/videos_ms/routes"
)

func main() {
	orm.Connection()
	routes.Start()
}
