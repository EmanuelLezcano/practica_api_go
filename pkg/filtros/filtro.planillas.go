package filtros

import "fmt"

type PlanillaFiltro struct {
	Recolectores_barrios_id uint
	Recolectores_id         uint
	Barrios_id              uint
	GetRecolecciones        bool
}

func (pf *PlanillaFiltro) Validate() (erro error) {

	if pf.Recolectores_barrios_id == 0 {
		if pf.Recolectores_id == 0 {
			return fmt.Errorf(ERROR_VALIDATE_IDS)
		}
	}
	return
}
