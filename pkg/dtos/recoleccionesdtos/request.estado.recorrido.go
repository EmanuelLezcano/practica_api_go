package recoleccionesdtos

import (
	"errors"

	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/entities"
)

type RequestEstadoRecorrido struct {
	Recorrido_id uint                         `json:"recorrido_id"`
	Estado       entities.EnumEstadoRecorrido `json:"estado"`
}

func (rer *RequestEstadoRecorrido) Validar() (erro error) {
	msj := "parametros no validos: "

	if rer.Recorrido_id <= 0 {
		erro = errors.New(msj + "debe enviar un id de recorrido valido")
		return
	}

	estado := rer.Estado
	erro = estado.Validate()
	return
}
