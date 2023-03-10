package entities

import "gorm.io/gorm"

type Departamentos struct {
	gorm.Model
	Departamento string
	Activo       bool
	ProvinciasId uint
	Provincia    Provincias   `json:"provincia" gorm:"foreignkey:provincias_id"`
	Localidads   []Localidads `json:"localidades"`
}
