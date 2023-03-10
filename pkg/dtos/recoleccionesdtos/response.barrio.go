package recoleccionesdtos

import "github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/entities"

type ResponseBarrios struct {
	Barrios []ResponseBarrio `json:"barrios"`
}

type ResponseBarrio struct {
	Id           uint   `json:"id"`
	Nombre       string `json:"nombre"`
	MunicipiosId uint   `json:"municipio_id"`
	Municipio    ResponseMunicipio
}

func (rb *ResponseBarrio) EntityBarrioToDtos(entityBarrio entities.Barrio) {
	rb.Id = entityBarrio.ID
	rb.Nombre = entityBarrio.Nombre
	rb.MunicipiosId = entityBarrio.MunicipiosId
}

func (rbs *ResponseBarrios) FromEntities(barrios []entities.Barrio) {
	for _, barr := range barrios {
		var tempResBarrio ResponseBarrio
		tempResBarrio.EntityBarrioToDtos(barr)
		rbs.Barrios = append(rbs.Barrios, tempResBarrio)
	}
}
