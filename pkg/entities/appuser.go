package entities

import (
	"gorm.io/gorm"
)

type Appuser struct {
	gorm.Model
	UsersAppId uint
	BarriosId  uint
	Barrio     Barrio `gorm:"foreignkey:barrios_id"`
}
