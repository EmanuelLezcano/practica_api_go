package entities

import "gorm.io/gorm"

type Provincias struct {
	gorm.Model
	Provincia     string
	Activo        bool
	Departamentos []Departamentos `json:"departamentos"`
}
