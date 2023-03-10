package recoleccionesdtos

import "errors"

type RequestRegistroAppUser struct {
	UserappId uint   `json:"userapp_id"`
	BarrioId  uint   `json:"barrio_id"`
	Rol       string `json:"rol"`
}

func (rra *RequestRegistroAppUser) Validar() (erro error) {
	msj := "parametros no validos: "
	if rra.UserappId <= 0 {
		erro = errors.New(msj + "debe enviar un id de usuario valido")
	}
	if rra.BarrioId <= 0 {
		erro = errors.New(msj + "debe enviar un id de barrio valido")
	}
	if len(rra.Rol) == 0 {
		erro = errors.New(msj + "debe enviar el rol del usuario")
	}
	return
}
