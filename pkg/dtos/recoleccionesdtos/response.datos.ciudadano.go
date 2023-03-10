package recoleccionesdtos

import (
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/entities"
	filtros "github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/filtros/recolecciones"
)

type ResponseAppuser struct {
	Id         uint
	UsersAppId uint
	BarriosId  uint
	Barrio     ResponseBarrio
}

func (rau *ResponseAppuser) EntityAppUserToDtos(filtro filtros.FiltroCargarDatos, entityAppUser entities.Appuser) {
	rau.Id = entityAppUser.ID
	rau.UsersAppId = entityAppUser.UsersAppId
	rau.BarriosId = entityAppUser.BarriosId
	if filtro.CargarBario {
		rau.Barrio.EntityBarrioToDtos(entityAppUser.Barrio)
	}
	if filtro.CargarMunicipio {
		rau.Barrio.Municipio.EntityMunicipioToDtos(entityAppUser.Barrio.Municipio)
	}
}
