package recolecciones

import (
	"context"
	"errors"
	"strings"

	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/internal/config"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/internal/storage"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/commons"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/domains/dbhelpers"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/dtos/recoleccionesdtos"
	"github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/entities"
	filtros "github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/filtros"
	filtros_recolecciones "github.com/Corrientes-Telecomunicaciones/api_go_recolector/pkg/filtros/recolecciones"
)

// definicion de interfaces
type RecoleccionesService interface {
	/*
		permite obtener la configuracion de botones por municipio
		recibe como parametro municipioId y responde una lista de configuraciones
	*/
	GetConfiguracionesByMunicipioIdService(municipioId uint) (listaConfiguraciones recoleccionesdtos.ResponseConfiguraciones, erro error)
	GetProvinciaService(request filtros.ProvinciaFiltro) (responseProvincias recoleccionesdtos.ResponseProvincias, erro error)
	GetDepartamentoService(request filtros.DepartamentoFiltro) (responseDepartamentos recoleccionesdtos.ResponseDepartamentos, erro error)
	GetLocalidadService(request filtros.LocalidadFiltro) (responseLocalidads recoleccionesdtos.ResponseLocalidads, erro error)
	GetMunicipioService(request filtros.MunicipioFiltro) (responseMunicipios recoleccionesdtos.ResponseMunicipios, erro error)
	GetBarrioService(request filtros.BarrioFiltro) (responseBarrios recoleccionesdtos.ResponseBarrios, erro error)
	GetPlanillasService(request filtros.PlanillaFiltro) (responsePlanillas recoleccionesdtos.ResponsePlanillaTrabajos, erro error)

	/*
		permite obtener el barrio y el municipio al que pertenece el ciudadano que utiliza la app
		recibe el users_app_id como parametro para obtener la informacion del servicio y retorna un objeto appusers
	*/
	GetAppUserByUserIdService(userAppId uint) (datosAppUser recoleccionesdtos.ResponseAppuser, rowAffected bool, erro error)

	// Cambia el estado de un recurrido (Enum:Esperando,Finalizado)
	PutEstadoRecorridoService(request recoleccionesdtos.RequestEstadoRecorrido) (status bool, erro error)

	// Cambia el estado de un recurrido (Enum:Esperando,Finalizado)
	PutEstadoPlanillaService(request recoleccionesdtos.RequestEstadoPlanilla) (status bool, erro error)

	/*
		permite crear la relacion entre el usiario registrado en login con el sistema de recolecciones
		recibe como parametro el rol, el id de usuario y el id del barrio
	*/
	PostRegistrarAppUsuarioService(usuarioApp recoleccionesdtos.RequestRegistroAppUser) (status bool, erro error)

	// Crear una recoleccion
	PostRecoleccionService(request recoleccionesdtos.RequestPostRecoleccion) (id uint, status bool, erro error)

	GetRecoleccionService(request filtros_recolecciones.GetRecoleccioneFiltro) (responseRecoleccione recoleccionesdtos.ResponseRecoleccione, rowAffected bool, erro error)

	PutRecoleccionService(request recoleccionesdtos.RequestPostRecoleccion) (status bool, erro error)

	// consultar por medio del dbHelpersRepository si existe un registro en una tabla de DB determinada
	CheckIfModelExists(tableName string, recolecciones_id uint) (res bool, erro error)

	// Guardar una imagen asociada a una recoleccion tanto en el storage como en la DB
	PostImagesRecoleccionService(request recoleccionesdtos.RequestImagesRecoleccionesDTO) (id uint, erro error)
}

// defino estructura
type recoleccionesService struct {
	repository          RecoleccionesRepository
	dbHelpersRepository dbhelpers.DbHelpersRepository
	store               storage.Storage // un objeto que implementa un metodo de almacenamiento
	commonsService      commons.Commons // otras funciones de ayuda
}

// defino un constructor
func NewRecoleccionesService(r RecoleccionesRepository, dbhr dbhelpers.DbHelpersRepository, st storage.Storage, c commons.Commons) RecoleccionesService {
	return &recoleccionesService{
		repository:          r,
		dbHelpersRepository: dbhr,
		store:               st,
		commonsService:      c,
	}
}

func (s *recoleccionesService) GetConfiguracionesByMunicipioIdService(municipioId uint) (listaConfiguraciones recoleccionesdtos.ResponseConfiguraciones, erro error) {
	entityConfiguraciones, err := s.repository.GetConfiguracionesByMunicipioIdRepository(municipioId)
	if err != nil {
		erro = errors.New(err.Error())
		return
	}

	for _, value := range entityConfiguraciones {
		var configuracionTemporal recoleccionesdtos.ResponseConfiguracion
		var botonTemporal recoleccionesdtos.ResponseBotone
		var tipoTemporal recoleccionesdtos.ResponseTipo
		configuracionTemporal.EntityConfiguracionToDtos(value)
		botonTemporal.EntityBotonToDtos(value.Boton)
		tipoTemporal.EntityTipoToDtos(value.Boton.Tipo)
		configuracionTemporal.Boton = botonTemporal
		configuracionTemporal.Boton.Tipo = tipoTemporal
		listaConfiguraciones.Configuraciones = append(listaConfiguraciones.Configuraciones, configuracionTemporal)
	}
	return
}

func (s *recoleccionesService) GetProvinciaService(request filtros.ProvinciaFiltro) (responseProvincias recoleccionesdtos.ResponseProvincias, erro error) {
	//valida request
	erro = request.Validate()
	if erro != nil {
		return
	}

	// requerir informacion al repositorio
	provincias, err := s.repository.GetProvinciaRepository(request)

	if err != nil {
		erro = errors.New(err.Error())
		return
	}
	responseProvincias.FromEntities(provincias)

	return
}

func (s *recoleccionesService) GetDepartamentoService(request filtros.DepartamentoFiltro) (responseDepartamentos recoleccionesdtos.ResponseDepartamentos, erro error) {
	//valida request
	erro = request.Validate()
	if erro != nil {
		return
	}

	// requerir informacion al repositorio
	departamentos, err := s.repository.GetDepartamentoRepository(request)

	if err != nil {
		erro = errors.New(err.Error())
		return
	}
	responseDepartamentos.FromEntities(departamentos)
	return
}

func (s *recoleccionesService) GetLocalidadService(request filtros.LocalidadFiltro) (responseLocalidads recoleccionesdtos.ResponseLocalidads, erro error) {

	//valida request
	erro = request.Validate()
	// requerir informacion al repositorio
	localidades, err := s.repository.GetLocalidadRepository(request)

	if err != nil {
		erro = errors.New(err.Error())
		return
	}

	responseLocalidads.FromEntities(localidades)

	return
}

func (s *recoleccionesService) GetMunicipioService(request filtros.MunicipioFiltro) (responseMunicipios recoleccionesdtos.ResponseMunicipios, erro error) {

	//valida request
	erro = request.Validate()
	// requerir informacion al repositorio
	municipios, err := s.repository.GetMunicipioRepository(request)

	if err != nil {
		erro = errors.New(err.Error())
		return
	}

	// to dto response
	responseMunicipios.FromEntities(municipios)

	return
}

func (s *recoleccionesService) GetAppUserByUserIdService(userAppId uint) (datosAppUser recoleccionesdtos.ResponseAppuser, rowAffected bool, erro error) {
	entityAppUser, rowAffected, err := s.repository.GetAppUserByUserIdRepository(userAppId)
	if err != nil {
		erro = errors.New(err.Error())
		return
	}
	if rowAffected {
		return
	}

	filtro := filtros_recolecciones.FiltroCargarDatos{
		CargarBario:     true,
		CargarMunicipio: true,
	}
	datosAppUser.EntityAppUserToDtos(filtro, entityAppUser)
	return
}

func (s *recoleccionesService) PostRegistrarAppUsuarioService(usuarioApp recoleccionesdtos.RequestRegistroAppUser) (status bool, erro error) {
	entityAppUser := entities.Appuser{
		UsersAppId: usuarioApp.UserappId,
		BarriosId:  usuarioApp.BarrioId,
	}
	if strings.ToUpper(usuarioApp.Rol) == "CIUDADANO" {
		err := s.repository.PostRegistrarAppUsuarioRepository(entityAppUser)
		if err != nil {
			status = false
			erro = errors.New(err.Error())
			return
		}
		status = true
	}
	return
}

func (s *recoleccionesService) GetBarrioService(request filtros.BarrioFiltro) (responseBarrios recoleccionesdtos.ResponseBarrios, erro error) {
	//valida request
	erro = request.Validate()
	// requerir informacion al repositorio
	barrios, err := s.repository.GetBarrioRepository(request)

	if err != nil {
		erro = errors.New(err.Error())
		return
	}

	// to dto response
	responseBarrios.FromEntities(barrios)

	return
}

func (s *recoleccionesService) GetPlanillasService(request filtros.PlanillaFiltro) (responsePlanillas recoleccionesdtos.ResponsePlanillaTrabajos, erro error) {
	//valida request
	erro = request.Validate()
	if erro != nil {
		return
	}

	recolectores_barrios, err := s.repository.GetRecolectoresBarriosRepository(request)
	if err != nil {
		erro = errors.New(err.Error())
		return
	}

	for _, Recolector_Barrio := range recolectores_barrios {

		auxFiltro := filtros.PlanillaFiltro{
			Recolectores_barrios_id: Recolector_Barrio.ID,
		}

		// requerir planillas al repositorio
		planillas, err := s.repository.GetPlanillasTrabajosRepository(auxFiltro)

		if err != nil {
			erro = errors.New(err.Error())
			return
		}

		for _, planilla := range planillas {

			var planillaResponse recoleccionesdtos.ResponsePlanillaTrabajo

			var filtroPlanillaRecorridos filtros.PlanillaRecorridoFiltro

			filtroPlanillaRecorridos.PlanillaId = planilla.ID

			// requerir planillas recorridos al repositorio
			planillaRecorridos, err := s.repository.GetPlanillasRecorridosRepository(filtroPlanillaRecorridos)

			if err != nil {
				erro = errors.New(err.Error())
				return
			}

			planillaResponse.FromRecorridos(planillaRecorridos)

			planillaResponse.FromEntity(planilla)

			responsePlanillas.LoadPlanilla(planillaResponse)

		}
	}

	return
}

func (s *recoleccionesService) PutEstadoRecorridoService(request recoleccionesdtos.RequestEstadoRecorrido) (status bool, erro error) {
	var entityRecorrido entities.Planillatrabajosrecorrido
	entityRecorrido.ID = request.Recorrido_id
	entityRecorrido.Estado = entities.EnumEstadoRecorrido(request.Estado)

	var recorridoFiltro = filtros.PlanillaRecorridoFiltro{
		Id: entityRecorrido.ID,
	}

	auxRecorridos, erro := s.repository.GetPlanillasRecorridosRepository(recorridoFiltro)

	for _, recorrido := range auxRecorridos {
		var EstadoRecorrido string

		switch entityRecorrido.Estado {
		case entities.EstadoRecorridoFinalizado:
			EstadoRecorrido = entities.Finalizado.String()

		case entities.EstadoRecorridoEsperando:
			EstadoRecorrido = entities.Asignado.String()
		}

		var recoleccionFiltro = filtros_recolecciones.GetRecoleccioneFiltro{
			Id: uint64(recorrido.ID),
		}

		auxRecolecciones, _, err := s.repository.GetRecoleccioneRepository(recoleccionFiltro)
		if err != nil {
			status = false
			erro = errors.New(err.Error())
			return
		}

		auxRecolecciones.Estado = entities.EnumRecoleccionEstado(EstadoRecorrido)

		err = s.repository.PutRecoleccionRepository(auxRecolecciones)

		if err != nil {
			status = false
			erro = errors.New(err.Error())
			return
		}
	}

	err := s.repository.PutEstadoRecorridoRepository(entityRecorrido)
	if erro != nil {
		status = false
		erro = errors.New(err.Error())
		return
	}
	status = true
	return
}

func (s *recoleccionesService) PutEstadoPlanillaService(request recoleccionesdtos.RequestEstadoPlanilla) (status bool, erro error) {
	var entityPlanilla entities.Planillatrabajo
	entityPlanilla.ID = request.Recolectores_barrios_id
	entityPlanilla.Estado = entities.EnumEstadoPlanillaTrabajo(request.Estado)

	if entityPlanilla.Estado == "Finalizado" {
		var filtroPlanillaRecorridos filtros.PlanillaRecorridoFiltro

		filtroPlanillaRecorridos.PlanillaId = entityPlanilla.ID

		// requerir planillas recorridos al repositorio
		planillaRecorridos, err := s.repository.GetPlanillasRecorridosRepository(filtroPlanillaRecorridos)

		if err != nil {
			erro = errors.New(err.Error())
			return
		}

		for _, recorrido := range planillaRecorridos {
			if recorrido.Estado != "Finalizado" {
				erro = errors.New(ERROR_UPDATE_ESTADO_PLANILLARECORRIDOS)
				return
			}
		}

	}

	err := s.repository.PutEstadoPlanillaRepository(entityPlanilla)
	if erro != nil {
		status = false
		erro = errors.New(err.Error())
		return
	}
	status = true
	return
}

func (s *recoleccionesService) PostRecoleccionService(request recoleccionesdtos.RequestPostRecoleccion) (id uint, status bool, erro error) {

	// validar datos de la request
	erro = request.Validate(false, false)
	if erro != nil {
		return 0, false, errors.New(erro.Error())
	}

	// solicitar repository
	id, erro = s.repository.PostRecoleccionRepository(request.ToEntity(false))

	if erro != nil {
		status = false
		erro = errors.New(erro.Error())
		return
	}
	status = true
	return
}

func (s *recoleccionesService) GetRecoleccionService(request filtros_recolecciones.GetRecoleccioneFiltro) (responseRecoleccione recoleccionesdtos.ResponseRecoleccione, rowAffected bool, erro error) {

	// validar datos de la request
	erro = request.Validate()
	if erro != nil {
		return
	}

	// get repository
	entity, rowAffected, err := s.repository.GetRecoleccioneRepository(request)
	if err != nil {
		erro = errors.New(err.Error())
		return
	}
	if rowAffected {
		return
	}

	responseRecoleccione.FromEntity(entity)
	return
}

func (s *recoleccionesService) PutRecoleccionService(request recoleccionesdtos.RequestPostRecoleccion) (status bool, erro error) {

	// validar datos de la request
	var OnlyEstado = false
	if request.ChangeEstado {
		OnlyEstado = true
		request = request.OnlyEstado()
	}
	erro = request.Validate(true, OnlyEstado)
	if erro != nil {
		return false, errors.New(erro.Error())
	}

	// solicitar repository
	erro = s.repository.PutRecoleccionRepository(request.ToEntity(true))
	if erro != nil {
		status = false
		erro = errors.New(erro.Error())
		return
	}
	status = true
	return
}

func (s *recoleccionesService) CheckIfModelExists(tableName string, recolecciones_id uint) (res bool, erro error) {
	return s.dbHelpersRepository.CheckIfModelsExists(tableName, recolecciones_id)
}

func (s *recoleccionesService) PostImagesRecoleccionService(request recoleccionesdtos.RequestImagesRecoleccionesDTO) (id uint, erro error) {

	// Validar la request
	erro = request.Validate()
	if erro != nil {
		return
	}

	// 1) Leer datos del archivo
	source := config.DIR_BASE + config.DIR_IMAGES_RECOLECCION  // carpeta temporal del archivo
	remoteFolder := config.DIR_BASE + config.DIR_IMAGES_REMOTE // carpeta destino en el remote storage

	// esta func le coloca una fecha al nombre del archivo
	fileData, fileName, fileType, erro := storage.LeerDatosArchivo(remoteFolder, source, request.FileName)
	// error para LeerDatosArchivo(...)
	if erro != nil {
		return
	}

	// se guarda la imagen en el storage. en caso de exito devuelve true y nil, o false y error en caso contrario
	res, erro := s.store.PutObject(context.Background(), fileData, fileName, fileType)
	if erro != nil {
		return
	}
	if res {
		// 2) Guardar en la DB
		// repositorio para guardar datos
		id, erro = s.repository.PostImagesRecoleccionRepository(request.ToImageRecoleccionEntity(false))
	}
	if erro != nil {
		return
	}

	// remover el archivo subido al remote storage de la carpeta temporal
	erro = s.commonsService.RemoveFile(source + "/" + request.FileName)
	if erro != nil {
		return
	}

	return
}
