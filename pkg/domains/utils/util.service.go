package utils

import (
	"fmt"
	"math"
	"strings"

	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/internal/logs"
	utils "github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/dtos/utils"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/entities"
	filtros "github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/filtros/utils"
)

type UtilService interface {
	CreateNotificacionService(notificacion entities.Notificacione) (erro error)
	CreateLogService(log entities.Log) (erro error)
	LogError(erro string, funcionalidad string)

	//CONFIGURACIONES
	GetConfiguracionService(filtro filtros.ConfiguracionFiltro) (configuracion utils.ResponseConfiguracion, erro error)
	CreateConfiguracionService(config utils.RequestConfiguracion) (id uint, erro error)
	FirstOrCreateConfiguracionService(nombre string, descripcion string, valor string) (key string, erro error)

	//Redondeo
	ToFixed(num float64, precision int) float64
}

func NewUtilService(r UtilRepository) UtilService {
	service := utilService{
		repository: r,
	}
	return &service
}

type utilService struct {
	repository UtilRepository
}

func (r *utilService) CreateNotificacionService(notificacion entities.Notificacione) (erro error) {
	return r.repository.CreateNotificacion(notificacion)

}

func (r *utilService) CreateLogService(log entities.Log) (erro error) {
	return r.repository.CreateLog(log)
}

func (r *utilService) LogError(erro string, funcionalidad string) {

	log := entities.Log{
		Tipo:          entities.Error,
		Mensaje:       erro,
		Funcionalidad: funcionalidad,
	}

	err := r.CreateLogService(log)

	if err != nil {
		mensaje := fmt.Sprintf("Crear Log: %s. %s", err.Error(), erro)
		logs.Error(mensaje)
	}
}

func (s *utilService) GetConfiguracionService(filtro filtros.ConfiguracionFiltro) (configuracion utils.ResponseConfiguracion, erro error) {

	response, erro := s.repository.GetConfiguracion(filtro)

	if erro != nil {
		return
	}

	configuracion.FromEntity(response)

	return
}

func (s *utilService) CreateConfiguracionService(config utils.RequestConfiguracion) (id uint, erro error) {

	erro = config.IsValid(false)

	if erro != nil {
		return
	}

	request := config.ToEntity(false)

	return s.repository.CreateConfiguracion(request)

}

func (s *utilService) FirstOrCreateConfiguracionService(nombre string, descripcion string, valor string) (key string, erro error) {

	if len(strings.TrimSpace(nombre)) < 1 || len(strings.TrimSpace(valor)) < 1 {
		erro = fmt.Errorf("el campo nombre o el campo valor es inválido")
		return
	}

	filtro := filtros.ConfiguracionFiltro{
		Nombre: nombre,
	}

	response, erro := s.GetConfiguracionService(filtro)

	if erro != nil || response.Id == 0 {

		configuracion := utils.RequestConfiguracion{
			Nombre:      nombre,
			Descripcion: descripcion,
			Valor:       valor,
		}
		_, erro = s.CreateConfiguracionService(configuracion)

		if erro != nil {
			return
		}

		response.Valor = valor
	}

	key = response.Valor

	return

}

func (s *utilService) ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
