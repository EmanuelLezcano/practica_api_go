package filtros

import (
	"fmt"
)

type ProvinciaFiltro struct {
	ProvinciaId uint
	GetAll      bool
}

func (pf *ProvinciaFiltro) Validate() (erro error) {

	if pf.ProvinciaId == 0 {
		return fmt.Errorf(ERROR_VALIDATE_ID)
	}
	return
}
