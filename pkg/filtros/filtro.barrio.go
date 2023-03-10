package filtros

import (
	"fmt"
)

type BarrioFiltro struct {
	MunicipioId uint
}

func (bf *BarrioFiltro) Validate() (erro error) {

	if bf.MunicipioId == 0 {
		return fmt.Errorf(ERROR_VALIDATE_ID)
	}

	return
}
