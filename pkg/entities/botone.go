package entities

import "gorm.io/gorm"

type Botone struct {
	gorm.Model
	Boton   string
	Titulo  string
	TiposId uint
	Tipo    Tipo `gorm:"foreignKey:tipos_id"`
}
