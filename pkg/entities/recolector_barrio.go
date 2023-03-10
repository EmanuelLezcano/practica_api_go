package entities

import (
	"gorm.io/gorm"
)

type Recolectores_barrio struct {
	gorm.Model
	Recolectores_id uint64 `json:"recolectores_id"`
	Barrios_id      uint64 `json:"barrios_id"`
}
