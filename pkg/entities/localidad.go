package entities

import "gorm.io/gorm"

type Localidads struct {
	gorm.Model
	DepartamentosId uint
	Localidad       string
	Activo          bool
	Departamentos   Departamentos `json:"departamento"`
	Municipios      Municipios    `json:"municipio"`
}
