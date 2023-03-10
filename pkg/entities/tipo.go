package entities

import "gorm.io/gorm"

type Tipo struct{
	gorm.Model
	Tipo string
	Descripcion string
}