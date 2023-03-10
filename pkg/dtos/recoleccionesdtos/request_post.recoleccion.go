package recoleccionesdtos

import (
	"errors"

	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/entities"
)

type RequestPostRecoleccion struct {
	Id           uint
	BotonesId    uint
	AppusersId   uint
	Descripcion  string
	Lat          string
	Lng          string
	Fake         bool
	Estado       string
	ChangeEstado bool
}

func (rpr *RequestPostRecoleccion) Validate(isUpdate bool, OnlyEstado bool) (erro error) {
	msj := "parametro no valido: "

	if !OnlyEstado {
		if isUpdate && rpr.Id < 1 {
			erro = errors.New(msj + "el id no es válido")
			return
		}

		if rpr.BotonesId <= 0 {
			erro = errors.New(msj + "debe enviar un id de boton")
			return
		}
		if rpr.AppusersId <= 0 {
			erro = errors.New(msj + "debe enviar un id de app user")
			return
		}
		if len(rpr.Lat) == 0 {
			erro = errors.New(msj + "debe enviar una latitud")
			return
		}
		if len(rpr.Lng) == 0 {
			erro = errors.New(msj + "debe enviar una longitud")
			return
		}
	}

	// crear un estado mediante la extension de tipo EnumRecoleccionEstado
	estado := entities.EnumRecoleccionEstado(rpr.Estado)
	// el tipo EnumRecoleccionEstado tiene asociado un metodo de validacion
	if estado.Validate() != nil {
		erro = errors.New(msj + "el estado debe ser uno válido")
		return
	}
	return
}

func (rpr *RequestPostRecoleccion) ToEntity(isUpdate bool) (e entities.Recoleccione) {
	// si es un  update se necesita el id de la entidad
	if isUpdate {
		e.ID = rpr.Id
	}
	e.BotonesId = rpr.BotonesId
	e.AppusersId = rpr.AppusersId
	e.Descripcion = rpr.Descripcion
	e.Lat = rpr.Lat
	e.Lng = rpr.Lng
	e.Fake = true
	e.Estado = entities.EnumRecoleccionEstado(rpr.Estado)
	return
}

func (rpr *RequestPostRecoleccion) OnlyEstado() (e RequestPostRecoleccion) {
	// si es un  OnlyEstado se necesita el id de la entidad y estado solamente
	var newRequest RequestPostRecoleccion
	newRequest.Id = rpr.Id
	newRequest.Estado = rpr.Estado
	e = newRequest
	return
}
