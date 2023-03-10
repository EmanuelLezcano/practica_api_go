package entities

import "gorm.io/gorm"

/* Relaciones
- una Recoleccionesimagen pertenece a una Recoleccione (belongs to)
*/

type Recoleccionesimagen struct {
	gorm.Model
	RecoleccionesId uint
	Recoleccion     Recoleccione `json:"recoleccion" gorm:"foreignKey:RecoleccionesId"` // belongs to Recoleccione
	FileName        string
}
