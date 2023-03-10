package recoleccionesdtos

import "github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/entities"

// estructura de response para endpoint de provincias
type ResponseProvincias struct {
	Provincias []ResponseProvincia `json:"provincias"`
}
type ResponseProvincia struct {
	Id            uint                   `json:"id"`
	Provincia     string                 `json:"provincia"`
	Activo        bool                   `json:"activo"`
	Departamentos []ResponseDepartamento `json:"departamentos"`
}

// Metodo asociado a ResponseProvincia
func (rp *ResponseProvincia) FromEntity(provincia entities.Provincias) {
	rp.Id = provincia.ID
	rp.Provincia = provincia.Provincia
	rp.Activo = provincia.Activo

	for _, d := range provincia.Departamentos {
		var tempResDepartamento ResponseDepartamento
		tempResDepartamento.FromEntity(d)
		rp.Departamentos = append(rp.Departamentos, tempResDepartamento)
	}
}

// Metodo asociado a ResponseProvincias
func (rps *ResponseProvincias) FromEntities(provincias []entities.Provincias) {
	for _, p := range provincias {
		var tempResProvincia ResponseProvincia
		tempResProvincia.FromEntity(p)
		rps.Provincias = append(rps.Provincias, tempResProvincia)
	}
}
