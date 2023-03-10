package recoleccionesdtos

import "github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/entities"

// estructura de response para endpoint de planillastrabajos
type ResponsePlanillaTrabajos struct {
	PlanillaTrabajos []ResponsePlanillaTrabajo `json:"planilla_trabajos"`
}
type ResponsePlanillaTrabajo struct {
	PlanillaTrabajo    entities.Planillatrabajo    `json:"planilla_trabajo"`
	PlanillaRecorridos []ResponsePlanillaRecorrido `json:"planilla_recorridos"`
}

type ResponsePlanillaRecorrido struct {
	ID                  uint                 `json:"id"`
	Planillatrabajos_id uint64               `json:"planillatrabajos_id"`
	RecoleccionesId     uint64               `json:"recolecciones_id"`
	Recoleccion         ResponseRecoleccione `json:"recoleccion"`
	Descripcion         string               `json:"descripcion"`
	Estadorecoleccion   int                  `json:"estadorecoleccion"`
	Lng                 string               `json:"lng"`
	Lat                 string               `json:"lat"`
	Estado              string               `json:"estado"`
}

func (rp *ResponsePlanillaTrabajo) FromEntity(e entities.Planillatrabajo) {
	rp.PlanillaTrabajo = e
}

func (rpt *ResponsePlanillaTrabajo) FromRecorridos(ArrayRecorridos []entities.Planillatrabajosrecorrido) {

	for _, e := range ArrayRecorridos {
		var rp ResponsePlanillaRecorrido
		rp.ID = e.ID
		rp.Descripcion = e.Descripcion
		rp.Estado = string(e.Estado)
		rp.Estadorecoleccion = e.Estadorecoleccion
		rp.Lat = e.Lat
		rp.Lng = e.Lng

		rp.Recoleccion.FromEntity(e.Recoleccion)

		rpt.PlanillaRecorridos = append(rpt.PlanillaRecorridos, rp)
	}

}

func (rps *ResponsePlanillaTrabajos) LoadPlanilla(responseLoad ResponsePlanillaTrabajo) {
	rps.PlanillaTrabajos = append(rps.PlanillaTrabajos, responseLoad)
}
