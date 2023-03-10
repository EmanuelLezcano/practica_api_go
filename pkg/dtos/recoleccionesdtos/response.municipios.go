package recoleccionesdtos

import "github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/entities"

// estructura de response para endpoint de departamentos
type ResponseMunicipios struct {
	Municipios []ResponseMunicipio `json:"municipios"`
}
type ResponseMunicipio struct {
	Id           uint             `json:"id"`
	Nombre       string           `json:"nombre"`
	Direccion    string           `json:"direccion"`
	Telefono     string           `json:"telefono"`
	LocalidadsId uint             `json:"localidads_id"`
	Barrios      []ResponseBarrio `json:"barrios"`
}

// Metodo asociado a ResponseMunicipio
func (rm *ResponseMunicipio) FromEntity(muni entities.Municipios) {
	rm.Id = muni.ID
	rm.Nombre = muni.Nombre
	rm.Direccion = muni.Direccion
	rm.Telefono = muni.Telefono
	rm.LocalidadsId = muni.LocalidadsId

	for _, barrio := range muni.Barrios {
		var tempResBarrio ResponseBarrio
		tempResBarrio.EntityBarrioToDtos(barrio)
		rm.Barrios = append(rm.Barrios, tempResBarrio)
	}
}

// Metodo asociado a ResponseMunicipios
func (rms *ResponseMunicipios) FromEntities(munis []entities.Municipios) {
	for _, m := range munis {
		var tempResMunicipio ResponseMunicipio
		tempResMunicipio.FromEntity(m)
		rms.Municipios = append(rms.Municipios, tempResMunicipio)
	}
}

func (rm *ResponseMunicipio) EntityMunicipioToDtos(entityMunicipio entities.Municipios) {
	rm.Id = entityMunicipio.ID
	rm.Nombre = entityMunicipio.Nombre
	rm.Direccion = entityMunicipio.Direccion
	rm.Telefono = entityMunicipio.Telefono
}
