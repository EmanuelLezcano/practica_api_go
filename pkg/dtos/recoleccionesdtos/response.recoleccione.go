package recoleccionesdtos

import (
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/entities"
)

type ResponseRecoleccione struct {
	Id          uint `json:"id"`
	BotonesId   uint
	AppusersId  uint
	Descripcion string
	Lat         string `json:"Latitud"`
	Lng         string `json:"Longitud"`
	Fake        bool
	Estado      string
	Creado      string `json:"Creado"`
	Modificado  string `json:"Modificado"`
}

func (rr *ResponseRecoleccione) FromEntity(entity entities.Recoleccione) {
	rr.Id = entity.ID
	rr.BotonesId = entity.BotonesId
	rr.AppusersId = entity.AppusersId
	rr.Descripcion = entity.Descripcion
	rr.Lat = entity.Lat
	rr.Lng = entity.Lng
	rr.Fake = entity.Fake
	rr.Estado = entity.Estado.String()
	rr.Creado = entity.CreatedAt.Format("2006-01-02 15:04:05")
	rr.Modificado = entity.UpdatedAt.Format("2006-01-02 15:04:05")
}
