package recoleccionesdtos

import "github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/entities"

type ResponseUbicacion struct {
	Id            uint                   `json:"id"`
	Provincia     string                 `json:"provincia"`
	Activo        bool                   `json:"activo"`
	Departamentos []ResponseDepartamento `json:"departamentos"`
}

func (ru *ResponseUbicacion) FromEntity(provincia entities.Provincias) {
	ru.Id = provincia.ID
	ru.Provincia = provincia.Provincia
	ru.Activo = provincia.Activo
	var responseDpto ResponseDepartamento
	for _, d := range provincia.Departamentos {
		responseDpto.FromEntity(d)
		ru.Departamentos = append(ru.Departamentos, responseDpto)
	}
}

// // Metodo asociado a ResponseProvincias
// func (rps *ResponseProvincias) FromEntities(provincias []entities.Provincias) {
// 	for _, p := range provincias {
// 		resp := FromEntity(p)
// 		rps.Provincias = append(rps.Provincias, resp)
// 	}
// }
