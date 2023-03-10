package filtros

import (
	"fmt"
)

type LocalidadFiltro struct {
	LocalidadId uint
	GetAll      bool
}

func (lf *LocalidadFiltro) Validate() (erro error) {

	if lf.LocalidadId == 0 {
		return fmt.Errorf(ERROR_VALIDATE_ID)
	}

	return
}
