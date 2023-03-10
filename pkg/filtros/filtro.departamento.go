package filtros

import (
	"fmt"
)

type DepartamentoFiltro struct {
	ProvinciaId    uint
	GetAll         bool
	DepartamentoId uint
}

func (df *DepartamentoFiltro) Validate() (erro error) {

	if df.ProvinciaId == 0 {
		return fmt.Errorf(ERROR_VALIDATE_ID)
	}

	if df.DepartamentoId == 0 {
		return fmt.Errorf(ERROR_VALIDATE_ID)
	}

	return
}
