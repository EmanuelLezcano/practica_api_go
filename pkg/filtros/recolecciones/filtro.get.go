package filtros

import "errors"

type GetRecoleccioneFiltro struct {
	Id uint64
}

func (grf *GetRecoleccioneFiltro) Validate() (erro error) {
	msj := "parametro no valido: "

	if grf.Id <= 0 {
		erro = errors.New(msj + "debe enviar un id de boton")
	}
	return
}
