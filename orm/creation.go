package orm

import "fmt"

// Create every table if not exist
func createDB() {
	db.AutoMigrate(&Video{}, &Musculos{}, &Grupo{}, &Objetivo{}, &Dificultad{}, &Equipamento{}, &Disciplina{})
	// fill default values for every category
	default_cat()
	fmt.Println("Database created")
}
