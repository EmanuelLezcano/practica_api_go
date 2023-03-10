package recoleccionesdtos

import "github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/entities"

// estructura de response para endpoint de departamentos
type ResponseLocalidads struct {
	Localidades []ResponseLocalidad `json:"localidades"`
}
type ResponseLocalidad struct {
	Id             uint              `json:"id"`
	Localidad      string            `json:"localidad"`
	Activo         bool              `json:"activo"`
	DepartamentoId uint              `json:"departamento_id"`
	Municipio      ResponseMunicipio `json:"municipio"`
}

// Metodo asociado a ResponseLocalidad
func (rl *ResponseLocalidad) FromEntity(localidad entities.Localidads) {
	rl.Id = localidad.ID
	rl.Localidad = localidad.Localidad
	rl.Activo = localidad.Activo
	rl.DepartamentoId = localidad.DepartamentosId
	// Se pasa la entidad municipio a type ResponseMunicipio
	var tempResMunicipio ResponseMunicipio
	tempResMunicipio.FromEntity(localidad.Municipios)
	rl.Municipio = tempResMunicipio
}

// Metodo asociado a ResponseLocalidads
func (rls *ResponseLocalidads) FromEntities(localidads []entities.Localidads) {
	for _, loc := range localidads {
		var tempResLocalidad ResponseLocalidad
		tempResLocalidad.FromEntity(loc)
		rls.Localidades = append(rls.Localidades, tempResLocalidad)
	}
}
