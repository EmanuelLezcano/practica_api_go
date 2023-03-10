package recoleccionesdtos

import (
	"errors"

	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/entities"
)

type RequestEstadoPlanilla struct {
	Recolectores_barrios_id uint                               `json:"recolectores_barrios_id"`
	Estado                  entities.EnumEstadoPlanillaTrabajo `json:"estado"`
}

func (rra *RequestEstadoPlanilla) Validar() (erro error) {
	msj := "parametros no validos: "

	if rra.Recolectores_barrios_id <= 0 {
		erro = errors.New(msj + "debe enviar un id de recolector_barrio valido")
		return
	}

	estado := rra.Estado
	erro = estado.Validate()

	return
}
