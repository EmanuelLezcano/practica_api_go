package filtros

import "fmt"

type PlanillaRecorridoFiltro struct {
	PlanillaId uint
	Id         uint
}

func (pf *PlanillaRecorridoFiltro) Validate() (erro error) {

	if pf.PlanillaId == 0 {
		return fmt.Errorf(ERROR_VALIDATE_ID)
	}
	return
}
