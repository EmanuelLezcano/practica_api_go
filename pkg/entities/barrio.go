package entities

import (
	"gorm.io/gorm"
)

type Barrio struct {
	gorm.Model
	Nombre       string
	MunicipiosId uint
	Municipio    Municipios `gorm:"foreignkey:municipios_id"`
}
