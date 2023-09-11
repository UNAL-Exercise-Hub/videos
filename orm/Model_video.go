package orm

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Titulo      string
	Link        string
	Imagen      string
	Duracion    int
	Series      int
	Musculos    []*Musculos    `gorm:"many2many:rel_musculos;"`
	Grupo       []*Grupo       `gorm:"many2many:rel_grupo;"`
	Objetivo    []*Objetivo    `gorm:"many2many:rel_objetivo;"`
	Dificultad  []*Dificultad  `gorm:"many2many:rel_dificultad;"`
	Equipamento []*Equipamento `gorm:"many2many:rel_equipamento;"`
	Disciplina  []*Disciplina  `gorm:"many2many:rel_disciplina;"`
}
