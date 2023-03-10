package entities

import (
	"errors"

	"gorm.io/gorm"
)

type Planillatrabajo struct {
	gorm.Model
	Recolectores_barrios_id uint64                    `json:"recolectores_barrios_id"`
	Descripcion             string                    `json:"descripcion"`
	Fecha_inicio            string                    `json:"fecha_inicio"`
	Fecha_fin               string                    `json:"fecha_fin"`
	Estado                  EnumEstadoPlanillaTrabajo `json:"estado"`
}

type EnumEstadoPlanillaTrabajo string

const (
	EstadoPlanillaTrabajoAsignado   EnumEstadoPlanillaTrabajo = "Asignado"
	EstadoPlanillaTrabajoIniciado   EnumEstadoPlanillaTrabajo = "Iniciado"
	EstadoPlanillaTrabajoFinalizado EnumEstadoPlanillaTrabajo = "Finalizado"
)

// validar si un string es alguno de los estados predefinidos
func (estado EnumEstadoPlanillaTrabajo) Validate() error {
	switch estado {
	case EstadoPlanillaTrabajoAsignado, EstadoPlanillaTrabajoIniciado, EstadoPlanillaTrabajoFinalizado:
		return nil
	}
	return errors.New("el tipo de estado de Planilla es inválido")
}
