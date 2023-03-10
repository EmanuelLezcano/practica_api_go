package entities

import (
	"errors"

	"gorm.io/gorm"
)

type Planillatrabajosrecorrido struct {
	gorm.Model
	Planillatrabajos_id uint64              `json:"planillatrabajos_id"`
	RecoleccionesId     uint64              `json:"recolecciones_id"`
	Recoleccion         Recoleccione        `json:"recoleccion" gorm:"foreignKey:RecoleccionesId"`
	Descripcion         string              `json:"descripcion"`
	Estadorecoleccion   int                 `json:"estadorecoleccion"`
	Lng                 string              `json:"lng"`
	Lat                 string              `json:"lat"`
	Estado              EnumEstadoRecorrido `json:"estado"`
}

type EnumEstadoRecorrido string

const (
	EstadoRecorridoEsperando  EnumEstadoRecorrido = "Esperando"
	EstadoRecorridoFinalizado EnumEstadoRecorrido = "Finalizado"
)

// validar si un string es alguno de los estados predefinidos
func (estado EnumEstadoRecorrido) Validate() error {
	switch estado {
	case EstadoRecorridoEsperando, EstadoRecorridoFinalizado:
		return nil
	}
	return errors.New("el tipo de estado de Recorrido es inválido")
}
