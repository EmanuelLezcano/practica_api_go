package filtros

import (
	"fmt"
)

type MunicipioFiltro struct {
	MunicipioId uint
	GetAll      bool
}

func (mf *MunicipioFiltro) Validate() (erro error) {

	if mf.MunicipioId == 0 {
		return fmt.Errorf(ERROR_VALIDATE_ID)
	}

	return
}
