package recoleccionesdtos

import "github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/entities"

// estructura de response para endpoint de departamentos
type ResponseDepartamentos struct {
	Departamentos []ResponseDepartamento `json:"departamentos"`
}
type ResponseDepartamento struct {
	Id           uint                `json:"id"`
	Departamento string              `json:"departamento"`
	Activo       bool                `json:"activo"`
	ProvinciaId  uint                `json:"provincia_id"`
	Localidades  []ResponseLocalidad `json:"localidades"`
}

// Metodo asociado a ResponseDepartamento
func (rd *ResponseDepartamento) FromEntity(departamento entities.Departamentos) {
	rd.Id = departamento.ID
	rd.Departamento = departamento.Departamento
	rd.Activo = departamento.Activo
	rd.ProvinciaId = departamento.ProvinciasId

	for _, loc := range departamento.Localidads {
		var tempResLocalidad ResponseLocalidad
		tempResLocalidad.FromEntity(loc)
		rd.Localidades = append(rd.Localidades, tempResLocalidad)
	}

}

// Metodo asociado a ResponseDepartamentos
func (rds *ResponseDepartamentos) FromEntities(departamentos []entities.Departamentos) {
	for _, ds := range departamentos {
		var tempResDepartamento ResponseDepartamento
		tempResDepartamento.FromEntity(ds)
		rds.Departamentos = append(rds.Departamentos, tempResDepartamento)
	}
}
