package recoleccionesdtos

import "github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/entities"

type ResponseConfiguraciones struct {
	Configuraciones []ResponseConfiguracion `json:"configuraciones"`
}

type ResponseConfiguracion struct {
	Id              uint   `json:"id"`
	MunicipiosId    uint   `json:"municipio_id"`
	BotonesId       uint   `json:"boton_id"`
	Nombre          string `json:"nombre"`
	Descripcion     string `json:"decripcion"`
	Valor           string `json:"valor"`
	DiasRecoleccion string
	Boton           ResponseBotone `json:"boton"`
}

type ResponseBotone struct {
	Id      uint         `json:"id"`
	Boton   string       `json:"boton"`
	Titulo  string       `json:"titulo"`
	TiposId uint         `json:"tipo_id"`
	Tipo    ResponseTipo `json:"tipo"`
}

type ResponseTipo struct {
	Id          uint   `json:"id"`
	Tipo        string `json:"tipo"`
	Descripcion string `json:"descripcion"`
}

func (rc *ResponseConfiguracion) EntityConfiguracionToDtos(entityConfiguracion entities.Configuracione) {
	rc.Id = entityConfiguracion.ID
	rc.MunicipiosId = entityConfiguracion.MunicipiosId
	rc.BotonesId = entityConfiguracion.BotonesId
	rc.Nombre = entityConfiguracion.Nombre
	rc.Descripcion = entityConfiguracion.Descripcion
	rc.Valor = entityConfiguracion.Valor
	rc.DiasRecoleccion = entityConfiguracion.DiasRecoleccion
}

func (rb *ResponseBotone) EntityBotonToDtos(entityBoiton entities.Botone) {
	rb.Id = entityBoiton.ID
	rb.Boton = entityBoiton.Boton
	rb.Titulo = entityBoiton.Titulo
	rb.TiposId = entityBoiton.TiposId
}

func (rt *ResponseTipo) EntityTipoToDtos(entityTipo entities.Tipo) {
	rt.Id = entityTipo.ID
	rt.Tipo = entityTipo.Tipo
	rt.Descripcion = entityTipo.Tipo
}
