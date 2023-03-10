package entities

import (
	"errors"

	"gorm.io/gorm"
)

/* Relaciones
- una recoleccion pertenece a un AppUser
- una recoleccion pertenece a un Botone
- una recoleccion tiene muchas Recoleccionesimagen
*/

type Recoleccione struct {
	gorm.Model
	BotonesId   uint    // se relaciona con tabla botones con la FK botones_id
	Boton       Botone  `json:"boton" gorm:"foreignKey:BotonesId"`
	AppusersId  uint    // se relaciona con tabla appusers con la FK appusers_id
	Appuser     Appuser `json:"appuser" gorm:"foreignKey:AppusersId"`
	Descripcion string
	Lat         string
	Lng         string
	Fake        bool
	Estado      EnumRecoleccionEstado
}

type EnumRecoleccionEstado string

const (
	Solicitado EnumRecoleccionEstado = "Solicitado"
	Anulado    EnumRecoleccionEstado = "Anulado"
	Cancelado  EnumRecoleccionEstado = "Cancelado"
	Asignado   EnumRecoleccionEstado = "Asignado"
	Finalizado EnumRecoleccionEstado = "Finalizado"
	Procesando EnumRecoleccionEstado = "Procesando"
)

// validar si un string es alguno de los estados predefinidos
func (estado EnumRecoleccionEstado) Validate() error {
	switch estado {
	case Solicitado, Asignado, Procesando, Finalizado, Anulado, Cancelado:
		return nil
	}
	return errors.New("el tipo de estado es inválido")
}

func (estado EnumRecoleccionEstado) String() string {
	switch estado {
	case Solicitado:
		return "Solicitado"
	case Asignado:
		return "Asignado"
	case Procesando:
		return "Procesando"
	case Finalizado:
		return "Finalizado"
	case Anulado:
		return "Anulado"
	case Cancelado:
		return "Cancelado"
	default:
		return "Atendido"
	}
}
