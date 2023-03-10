package entities

import "gorm.io/gorm"

type Municipios struct {
	gorm.Model
	LocalidadsId uint
	Nombre       string
	Direccion    string
	Telefono     string
	Barrios      []Barrio `json:"barrios"`
}
