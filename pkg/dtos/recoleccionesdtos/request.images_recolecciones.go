package recoleccionesdtos

import (
	"errors"

	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/entities"
)

type RequestImagesRecoleccionesDTO struct {
	RecoleccionesId uint   ` form:"RecoleccionesId"`
	FileName        string `form:"FileName"`
	Path            string
}

func (rir *RequestImagesRecoleccionesDTO) Validate() (erro error) {
	msj := "parametros no validos: "

	if rir.RecoleccionesId <= 0 {
		erro = errors.New(msj + "debe enviar un id de recoleccion valido")
		return
	}

	if len(rir.FileName) == 0 {
		erro = errors.New(msj + "debe enviar un nombre de imagen valido")
		return
	}

	if len(rir.Path) == 0 {
		erro = errors.New(msj + "la ruta de la imagen esta vacia")
		return
	}

	return
}

func (rir *RequestImagesRecoleccionesDTO) ToImageRecoleccionEntity(isUpdate bool) (e entities.Recoleccionesimagen) {
	// si es un  update se necesita el id de la entidad
	// if isUpdate {
	// 	e.ID = rir.Id
	// }
	e.RecoleccionesId = rir.RecoleccionesId
	e.FileName = rir.FileName
	return
}
